package edb

import "sync"

type TableColumn struct {

  StringTable map[string]int32

  string_id_m *sync.Mutex;
  val_string_id_lookup map[int32]string
}

func newTableColumn() *TableColumn {
  tc := TableColumn{}
  tc.StringTable = make(map[string]int32)
  tc.val_string_id_lookup = make(map[int32]string)
  tc.string_id_m = &sync.Mutex{}

  return &tc
}

func (tc *TableColumn) get_val_id(name string) int32 {

  id, ok := tc.StringTable[name]

  if ok {
    return int32(id);
  }


  tc.string_id_m.Lock();
  tc.StringTable[name] = int32(len(tc.StringTable));
  tc.val_string_id_lookup[tc.StringTable[name]] = name;
  tc.string_id_m.Unlock();
  return tc.StringTable[name];
}


func (tc *TableColumn) get_string_for_val(id int32) string {
  val, _ := tc.val_string_id_lookup[id];
  return val
}

