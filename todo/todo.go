package todo

import (
  "fmt"
  "encoding/json"
  "io/ioutil"
)

type Item struct {
  Text string
  Priority int
}

func (i *Item) SetPriority(pri int) {
  switch pri {
  case 1:
    i.Priority = 1
  case 3:
    i.Priority = 3
  default:
    i.Priority = 2
  }
}

func (i *Item) PrettyP() string {
  if i.Priority == 1 {
    return "(1)"
  }
  if i.Priority == 3 {
    return "(3)"
  }
  return ""
}

func SaveItems(filename string, items []Item) error {
  b, err := json.Marshal(items)
  if err != nil {
    return err
  }
  err = ioutil.WriteFile(filename, b, 0644) // last is chmod
  if err != nil {
    return err
  }
  fmt.Println(string(b))
  return nil
}

func ReadItems(filename string) ([]Item, error) {
  b, err := ioutil.ReadFile(filename)
  if err != nil {
    return []Item{}, err
  }
  var items []Item
  if err := json.Unmarshal(b, &items); err != nil {
    return []Item{}, err
  }

  return items, nil
}
