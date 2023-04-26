//@author Stanislav Polaniev <spolanyev@gmail.com>

package main

type CommandInterface interface {
	Execute(params *map[string]interface{}) map[string]interface{}
}
