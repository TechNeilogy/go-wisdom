package configoopw

import "errors"

type ConfigBuilder struct {
	err   *error
	name  string
	isErr bool
	count *int
}

type Config struct {
	Name  string
	IsErr bool
	Count int
}

func New() *ConfigBuilder {
	return &ConfigBuilder{}
}

func (cb *ConfigBuilder) Name(name string) *ConfigBuilder {
	if cb.err != nil {
		return cb
	}
	cb.name = name
	return cb
}

func (cb *ConfigBuilder) IsErr(isErr bool) *ConfigBuilder {
	if cb.err != nil {
		return cb
	}
	cb.isErr = isErr
	if isErr {
		err := errors.New("IsErr")
		cb.err = &err
	}
	return cb
}

func (cb *ConfigBuilder) Count(count int) *ConfigBuilder {
	if cb.err != nil {
		return cb
	}
	cb.count = &count
	return cb
}

func (cb *ConfigBuilder) Build() (*Config, error) {
	if cb.err != nil {
		return nil, *cb.err
	}
	c := &Config{
		Name:  cb.name,
		IsErr: cb.isErr,
	}
	if cb.count != nil {
		c.Count = *(cb.count)
	}
	return c, nil
}
