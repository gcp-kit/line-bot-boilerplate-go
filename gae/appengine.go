package gae

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	cloudtasks "cloud.google.com/go/cloudtasks/apiv2"
	"github.com/gcp-kit/line-bot-boilerplate-go/core"
	"github.com/gcp-kit/line-bot-boilerplate-go/util"
	"github.com/line/line-bot-sdk-go/linebot"
	"golang.org/x/xerrors"
	"google.golang.org/genproto/googleapis/cloud/tasks/v2"
)

// Props ...
type Props struct {
	QueuePath   string
	RelativeURI string

	client *cloudtasks.Client
	secret string
}

// NewProps - constructor
func NewProps(client *cloudtasks.Client) *Props {
	return &Props{
		client: client,
	}
}

// SetSecret - setter
func (p *Props) SetSecret(secret string) {
	p.secret = secret
}

func (p *Props) createTask(ctx context.Context, data []byte) error {
	req := &tasks.CreateTaskRequest{
		Parent: p.QueuePath,
		Task: &tasks.Task{
			MessageType: &tasks.Task_AppEngineHttpRequest{
				AppEngineHttpRequest: &tasks.AppEngineHttpRequest{
					HttpMethod:  tasks.HttpMethod_POST,
					RelativeUri: p.RelativeURI,
				},
			},
		},
	}

	req.Task.GetAppEngineHttpRequest().Body = data

	if _, err := p.client.CreateTask(ctx, req); err != nil {
		return xerrors.Errorf("failed to create tasks: %w", err)
	}

	return nil
}

// LineWebHook ...
func (p *Props) LineWebHook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, "NG", http.StatusInternalServerError)
		return
	}

	if !util.ValidateSignature(p.secret, r.Header.Get("X-Line-Signature"), body) {
		log.Println("invalid signature.")
		http.Error(w, "NG", http.StatusBadRequest)
		return
	}

	if err = p.createTask(r.Context(), body); err != nil {
		log.Printf("error: %+v\n", err)
		http.Error(w, "NG", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, http.StatusText(http.StatusOK))
}

// ParentFunctions ...
func (p *Props) ParentFunctions(ctx context.Context, body []byte) error {
	var (
		wg     sync.WaitGroup
		events = util.ParseEvents(body)
	)

	for _, event := range events {
		// nolint
		wg.Add(1)
		go func(ev *linebot.Event) {
			defer wg.Done()
			data, err := ev.MarshalJSON()
			if err != nil {
				return
			}
			if err = p.createTask(ctx, data); err != nil {
				log.Printf("error: %+v\n", err)
				return
			}
		}(event)
	}
	wg.Wait()
	return nil
}

// ChildFunctions ...
func (p *Props) ChildFunctions(ctx context.Context, op *core.Operation, body []byte) error {
	var event *linebot.Event
	if err := event.UnmarshalJSON(body); err != nil {
		return err
	}

	if err := op.Switcher(ctx, event); err != nil {
		if len(op.ErrMessage) > 0 {
			if er := op.SendReplyMessage(event.ReplyToken, op.ErrMessage); er != nil {
				log.Println(er)
			}
		}
		return err
	}
	return nil
}
