package testing

import (
	"context"
	"fmt"
	"investment/utility"

	"testing"

	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/schema"
)

func TestBase(t *testing.T) {
	template := prompt.FromMessages(
		schema.GoTemplate,
		schema.SystemMessage("你是一个{{.role}}。你需要用{{.style}}的语气回答问题。你的目标是帮助程序员保持积极乐观的心态，提供技术建议的同时也要关注他们的心理健康。"),
		schema.MessagesPlaceholder("chat_history", true),
		schema.UserMessage("问题: {{.question}}"),
	)

	msg, err := template.Format(context.Background(), map[string]any{
		"role":     "程序员鼓励师",
		"style":    "积极、温暖且专业",
		"question": "我的代码一直报错，感觉好沮丧，该怎么办？",
		"chat_history": []*schema.Message{
			schema.UserMessage("你好"),
			schema.AssistantMessage("嘿！我是你的程序员鼓励师！记住，每个优秀的程序员都是从 Debug 中成长起来的。有什么我可以帮你的吗？", nil),
			schema.UserMessage("我觉得自己写的代码太烂了"),
			schema.AssistantMessage("每个程序员都经历过这个阶段！重要的是你在不断学习和进步。让我们一起看看代码，我相信通过重构和优化，它会变得更好。记住，Rome wasn't built in a day，代码质量是通过持续改进来提升的。", nil),
		},
	})
	if err != nil {
		panic(err)
	}
	llm := utility.GetQwenChatModel()

	resp, err := llm.Generate(context.Background(), msg, model.WithTemperature(0.7))
	if err != nil {
		panic(err)
	}
	out := utility.Jsonout(resp)
	fmt.Println(out)
}
