package micro

import (
	"fmt"
)

type GiftKey struct {
	UserID string
	PostIDs []string
}

func (k GiftKey) String() string { return fmt.Sprint(k.UserID, k.PostIDs)  }

func (k GiftKey) Raw() interface{} { return k}

//type BoardKey struct {
//	ID    int64
//	isDev bool
//}
//
//func (k BoardKey) String() string { return fmt.Sprint(k.ID, k.isDev) }
//
//func (k BoardKey) Raw() interface{} { return k }
