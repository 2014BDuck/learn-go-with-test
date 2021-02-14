// @Author: 2014BDuck
// @Date: 2021/2/14

package dictionary

type Dictionary map[string]string

const (
	ErrorNotFound   = dictionaryErr("could not find the word you were looking for")
	ErrorWordExists = dictionaryErr("word exists")
)

type dictionaryErr string

func (e dictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(k string) (v string, err error) {
	v, ok := d[k]
	if !ok {
		return "", ErrorNotFound
	}
	return v, nil
}

func (d Dictionary) Add(k, v string) error {
	_, err := d.Search(k)
	switch err {
	case ErrorNotFound:
		d[k] = v
	case nil:
		return ErrorWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(k, v string) error {
	_, err := d.Search(k)
	switch err {
	case ErrorNotFound:
		return err
	case nil:
		d[k] = v
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(k string) error {
	_, err := d.Search(k)
	switch err {
	case ErrorNotFound:
		return err
	case nil:
		delete(d, k)
	default:
		return err
	}
	return nil
}
