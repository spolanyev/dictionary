//@author Stanislav Polaniev <spolanyev@gmail.com>

package command

import "dictionary/dto"

type CommandInterface interface {
	Execute(params map[string]interface{}) dto.ResponseInterface
}
