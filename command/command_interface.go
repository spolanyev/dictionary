//@author Stanislav Polaniev <spolanyev@gmail.com>

package command

import "dictionary/dto"

type CommandInterface interface {
	Execute(payload dto.RequestInterface) dto.ResponseInterface
	GetName() CommandName
}
