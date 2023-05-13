package configfuncw

import "errors"

type ConfigBuilder struct {
	name  string
	isErr bool
	count *int
}

func (cb *ConfigBuilder) Build() (*Config, error) {
	c := &Config{}
	c.Name = cb.name
	c.IsErr = cb.isErr
	if cb.count != nil {
		c.Count = *(cb.count)
	}
	return c, nil
}

type Config struct {
	Name  string
	IsErr bool
	Count int
}

type Option func(cb *ConfigBuilder) error

func Name(name string) Option {
	return func(cb *ConfigBuilder) error {
		cb.name = name
		return nil
	}
}

func IsErr(isErr bool) Option {
	return func(cb *ConfigBuilder) error {
		cb.isErr = isErr
		if isErr {
			return errors.New("IsErr")
		} else {
			return nil
		}
	}
}

func Count(count int) Option {
	return func(cb *ConfigBuilder) error {
		cb.count = &count
		return nil
	}
}

func Make(options ...Option) (*Config, error) {
	cb := &ConfigBuilder{}
	for _, o := range options {
		err := o(cb)
		if err != nil {
			return nil, err
		}
	}
	return cb.Build()
}
