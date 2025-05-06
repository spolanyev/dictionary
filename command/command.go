//@author Stanislav Polaniev <spolanyev@gmail.com>

package command

import "dictionary/dto"

type Command interface {
	Execute(payload dto.Request) dto.Response
	GetName() Name
}
