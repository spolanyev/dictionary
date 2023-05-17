//@author Stanislav Polaniev <spolanyev@gmail.com>

package command

type Interface interface {
	Execute(params map[string]interface{}) map[string]interface{}
}
