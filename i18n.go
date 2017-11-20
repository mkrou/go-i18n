package i18n

import (
	"errors"
	"strings"
	"fmt"
)

const separator = "."

type i18n struct {
	hash    map[string]hash
	current []string
}
type hash map[string]string
type Map map[string]interface{}

var internal = i18n{hash: make(map[string]hash)}

func (i i18n) T(key string, args ... interface{}) string {
	value := i.find(key)
	if len(args) > 0 {
		return fmt.Sprintf(value, args...)
	}
	return value
}
func (i i18n) ErrT(key string, args ...interface{}) error {
	return errors.New(i.T(key, args...))
}
func (i *i18n) Current(languages []string) error {
	for _, lang := range languages {
		if _, ok := i.hash[lang]; !ok {
			return errors.New(fmt.Sprintf("Language [%s] is undefined.", lang))
		}
	}
	i.current = languages
	return nil
}
func (i *i18n) AddLanguage(lang string, m Map) error {
	hash, err := m.genHash()
	if err != nil {
		return err
	}
	i.hash[lang] = hash
	return nil
}
func (i i18n) find(key string) string {
	for _, lang := range i.current {
		if m, ok := i.hash[lang]; ok {
			if value, ok := m[key]; ok {
				return value
			}
		}
	}
	return key
}
func (first *hash) join(second hash) {
	for key, value := range second {
		(*first)[key] = value
	}
}
func (m Map) genHash(prefix ... string) (hash, error) {
	var result hash = make(hash)
	for key, value := range m {
		switch x := value.(type) {
		case string:
			result[strings.Join(append(prefix, key), separator)] = x
		case Map:
			if r, err := x.genHash(append(prefix, key)...); err != nil {
				return nil, err
			} else {
				result.join(r)
			}
		default:
			return nil, errors.New(fmt.Sprintf("Type %T unsupported. Please use string or %T", x, Map{}))
		}
	}
	return result, nil
}
func AddLanguage(lang string, m Map) error {
	err := internal.AddLanguage(lang, m)
	if err != nil {
		return fmt.Errorf("Can't add language [%s]: %s.", lang, err)
	}
	return nil
}
func Current(languages []string) error {
	return internal.Current(languages)
}
func T(key string, args ... interface{}) string {
	return internal.T(key, args...)
}
func ErrT(key string, args ... interface{}) error {
	return internal.ErrT(key, args...)
}
