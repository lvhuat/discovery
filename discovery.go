package sd

import (
	"fmt"
)

// Discovery 定义了服务发现的接口
type Discovery interface {
	FindService(string) ([]string, []string, error)
}

// discovery 具体实现
type discovery struct {
	finders []func(string) ([]string, []string, error)
}

var _ Discovery = (*discovery)(nil)

var ErrUnkownService = fmt.Errorf("not avaliable discovery")

// Discovery 发现服务
func (discovery *discovery) FindService(serviceName string) ([]string, []string, error) {
	for _, find := range discovery.finders {
		endpoints, nodes, err := find(serviceName)
		if err != nil {
			return nil, nil, err
		}

		return endpoints, nodes, nil
	}

	return nil, nil, ErrUnkownService
}

func New(finders ...func(string) ([]string, []string, error)) Discovery {
	return &discovery{
		finders: finders,
	}
}
