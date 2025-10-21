package options

import "fmt"

type Options struct {
	Headers map[string]string
	Timeout int
}

func NewOptions() *Options {
	return &Options{
		Headers: map[string]string{},
		Timeout: 30,
	}
}

func (o *Options) getHeader() map[string]string {
	copy := make(map[string]string)

	for k, v := range o.Headers {
		copy[k] = v
	}

	return copy
}

func (o *Options) setHeaders(headers map[string]string) error {
	allow := map[string]bool{
		"Content-type":  true,
		"Accept":        true,
		"Authorization": true,
		"User-Agent":    true,
	}

	for k, v := range headers {
		if !allow[k] {
			fmt.Errorf("header %q unallowed", k)
		}
		o.Headers[k] = v
	}
	return nil
}
