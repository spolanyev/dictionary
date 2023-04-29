//@author Stanislav Polaniev <spolanyev@gmail.com>

package command

type CommandInterface interface {
	Execute(params map[string]interface{}) map[string]interface{}
}
