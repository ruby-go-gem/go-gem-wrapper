// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated
// Code generated by ruby_h_to_go. DO NOT EDIT.

package ruby

/*
#include "ruby.h"
*/
import "C"

import (
	"unsafe"
)

// StCheckForSizeofStIndexT is a type for passing `C.st_check_for_sizeof_st_index_t` in and out of package
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/st.h
type StCheckForSizeofStIndexT C.st_check_for_sizeof_st_index_t

// StCompareFunc is a type for passing `C.st_compare_func` in and out of package
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/st.h
type StCompareFunc C.st_compare_func

// StDataT is a type for passing `C.st_data_t` in and out of package
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/st.h
type StDataT C.st_data_t

// StForeachCallbackFunc is a type for passing `C.st_foreach_callback_func` in and out of package
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/st.h
type StForeachCallbackFunc C.st_foreach_callback_func

// StForeachCheckCallbackFunc is a type for passing `C.st_foreach_check_callback_func` in and out of package
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/st.h
type StForeachCheckCallbackFunc C.st_foreach_check_callback_func

// StHashFunc is a type for passing `C.st_hash_func` in and out of package
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/st.h
type StHashFunc C.st_hash_func

// StIndexT is a type for passing `C.st_index_t` in and out of package
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/st.h
type StIndexT C.st_index_t

// StTable is a type for passing `C.st_table` in and out of package
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/st.h
type StTable C.st_table

// StTableEntry is a type for passing `C.st_table_entry` in and out of package
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/st.h
type StTableEntry C.st_table_entry

// StUpdateCallbackFunc is a type for passing `C.st_update_callback_func` in and out of package
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/st.h
type StUpdateCallbackFunc C.st_update_callback_func

// RbHashBulkInsertIntoStTable calls `rb_hash_bulk_insert_into_st_table` in C
//
// Original definition is following
//
//	void rb_hash_bulk_insert_into_st_table(long, const VALUE *, VALUE)
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/st.h
func RbHashBulkInsertIntoStTable(arg1 Long, arg2 *VALUE, arg3 VALUE) {
	var cArg2 C.VALUE
	C.rb_hash_bulk_insert_into_st_table(C.long(arg1), &cArg2, C.VALUE(arg3))
	*arg2 = VALUE(cArg2)
}

// RbStAddDirect calls `rb_st_add_direct` in C
//
// Original definition is following
//
//	void rb_st_add_direct(st_table *, st_data_t, st_data_t)
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/st.h
func RbStAddDirect(arg1 *StTable, arg2 StDataT, arg3 StDataT) {
	var cArg1 C.st_table
	C.rb_st_add_direct(&cArg1, C.st_data_t(arg2), C.st_data_t(arg3))
	*arg1 = StTable(cArg1)
}

// RbStCleanupSafe calls `rb_st_cleanup_safe` in C
//
// Original definition is following
//
//	void rb_st_cleanup_safe(st_table *, st_data_t)
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/st.h
func RbStCleanupSafe(arg1 *StTable, arg2 StDataT) {
	var cArg1 C.st_table
	C.rb_st_cleanup_safe(&cArg1, C.st_data_t(arg2))
	*arg1 = StTable(cArg1)
}

// RbStClear calls `rb_st_clear` in C
//
// Original definition is following
//
//	void rb_st_clear(st_table *)
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/st.h
func RbStClear(arg1 *StTable) {
	var cArg1 C.st_table
	C.rb_st_clear(&cArg1)
	*arg1 = StTable(cArg1)
}

// RbStCopy calls `rb_st_copy` in C
//
// Original definition is following
//
//	st_table *rb_st_copy(st_table *)
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/st.h
func RbStCopy(arg1 *StTable) *StTable {
	var cArg1 C.st_table
	ret := (*StTable)(C.rb_st_copy(&cArg1))
	*arg1 = StTable(cArg1)
	return ret
}

// RbStDelete calls `rb_st_delete` in C
//
// Original definition is following
//
//	int rb_st_delete(st_table *, st_data_t *, st_data_t *); \/* returns 0:notfound 1:deleted *\/
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/st.h
func RbStDelete(arg1 *StTable, arg2 *StDataT, arg3 *StDataT) int {
	var cArg1 C.st_table
	var cArg2 C.st_data_t
	var cArg3 C.st_data_t
	ret := int(C.rb_st_delete(&cArg1, &cArg2, &cArg3))
	*arg1 = StTable(cArg1)
	*arg2 = StDataT(cArg2)
	*arg3 = StDataT(cArg3)
	return ret
}

// RbStDeleteSafe calls `rb_st_delete_safe` in C
//
// Original definition is following
//
//	int rb_st_delete_safe(st_table *, st_data_t *, st_data_t *, st_data_t)
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/st.h
func RbStDeleteSafe(arg1 *StTable, arg2 *StDataT, arg3 *StDataT, arg4 StDataT) int {
	var cArg1 C.st_table
	var cArg2 C.st_data_t
	var cArg3 C.st_data_t
	ret := int(C.rb_st_delete_safe(&cArg1, &cArg2, &cArg3, C.st_data_t(arg4)))
	*arg1 = StTable(cArg1)
	*arg2 = StDataT(cArg2)
	*arg3 = StDataT(cArg3)
	return ret
}

// RbStForeach calls `rb_st_foreach` in C
//
// Original definition is following
//
//	int rb_st_foreach(st_table *, st_foreach_callback_func *, st_data_t)
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/st.h
func RbStForeach(arg1 *StTable, arg2 *StForeachCallbackFunc, arg3 StDataT) int {
	var cArg1 C.st_table
	var cArg2 C.st_foreach_callback_func
	ret := int(C.rb_st_foreach(&cArg1, &cArg2, C.st_data_t(arg3)))
	*arg1 = StTable(cArg1)
	*arg2 = StForeachCallbackFunc(cArg2)
	return ret
}

// RbStForeachCheck calls `rb_st_foreach_check` in C
//
// Original definition is following
//
//	int rb_st_foreach_check(st_table *, st_foreach_check_callback_func *, st_data_t, st_data_t)
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/st.h
func RbStForeachCheck(arg1 *StTable, arg2 *StForeachCheckCallbackFunc, arg3 StDataT, arg4 StDataT) int {
	var cArg1 C.st_table
	var cArg2 C.st_foreach_check_callback_func
	ret := int(C.rb_st_foreach_check(&cArg1, &cArg2, C.st_data_t(arg3), C.st_data_t(arg4)))
	*arg1 = StTable(cArg1)
	*arg2 = StForeachCheckCallbackFunc(cArg2)
	return ret
}

// RbStForeachWithReplace calls `rb_st_foreach_with_replace` in C
//
// Original definition is following
//
//	int rb_st_foreach_with_replace(st_table *tab, st_foreach_check_callback_func *func, st_update_callback_func *replace, st_data_t arg)
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/st.h
func RbStForeachWithReplace(tab *StTable, fun *StForeachCheckCallbackFunc, replace *StUpdateCallbackFunc, arg StDataT) int {
	var cTab C.st_table
	var cFun C.st_foreach_check_callback_func
	var cReplace C.st_update_callback_func
	ret := int(C.rb_st_foreach_with_replace(&cTab, &cFun, &cReplace, C.st_data_t(arg)))
	*tab = StTable(cTab)
	*fun = StForeachCheckCallbackFunc(cFun)
	*replace = StUpdateCallbackFunc(cReplace)
	return ret
}

// RbStFreeTable calls `rb_st_free_table` in C
//
// Original definition is following
//
//	void rb_st_free_table(st_table *)
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/st.h
func RbStFreeTable(arg1 *StTable) {
	var cArg1 C.st_table
	C.rb_st_free_table(&cArg1)
	*arg1 = StTable(cArg1)
}

// RbStGetKey calls `rb_st_get_key` in C
//
// Original definition is following
//
//	int rb_st_get_key(st_table *, st_data_t, st_data_t *)
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/st.h
func RbStGetKey(arg1 *StTable, arg2 StDataT, arg3 *StDataT) int {
	var cArg1 C.st_table
	var cArg3 C.st_data_t
	ret := int(C.rb_st_get_key(&cArg1, C.st_data_t(arg2), &cArg3))
	*arg1 = StTable(cArg1)
	*arg3 = StDataT(cArg3)
	return ret
}

// RbStInitExistingTableWithSize calls `rb_st_init_existing_table_with_size` in C
//
// Original definition is following
//
//	st_table *rb_st_init_existing_table_with_size(st_table *tab, const struct st_hash_type *type, st_index_t size)
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/st.h
func RbStInitExistingTableWithSize(tab *StTable, r *StHashType, size StIndexT) *StTable {
	var cTab C.st_table
	var cR C.st_hash_type
	ret := *StTable(C.rb_st_init_existing_table_with_size(&cTab, &cR, C.st_index_t(size)))
	*tab = StTable(cTab)
	*r = StHashType(cR)
	return ret
}

// RbStInitNumtable calls `rb_st_init_numtable` in C
//
// Original definition is following
//
//	st_table *rb_st_init_numtable(void)
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/st.h
func RbStInitNumtable() *StTable {
	ret := (*StTable)(C.rb_st_init_numtable())
	return ret
}

// RbStInitNumtableWithSize calls `rb_st_init_numtable_with_size` in C
//
// Original definition is following
//
//	st_table *rb_st_init_numtable_with_size(st_index_t)
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/st.h
func RbStInitNumtableWithSize(arg1 StIndexT) *StTable {
	ret := (*StTable)(C.rb_st_init_numtable_with_size(C.st_index_t(arg1)))
	return ret
}

// RbStInitStrcasetable calls `rb_st_init_strcasetable` in C
//
// Original definition is following
//
//	st_table *rb_st_init_strcasetable(void)
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/st.h
func RbStInitStrcasetable() *StTable {
	ret := (*StTable)(C.rb_st_init_strcasetable())
	return ret
}

// RbStInitStrcasetableWithSize calls `rb_st_init_strcasetable_with_size` in C
//
// Original definition is following
//
//	st_table *rb_st_init_strcasetable_with_size(st_index_t)
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/st.h
func RbStInitStrcasetableWithSize(arg1 StIndexT) *StTable {
	ret := (*StTable)(C.rb_st_init_strcasetable_with_size(C.st_index_t(arg1)))
	return ret
}

// RbStInitStrtable calls `rb_st_init_strtable` in C
//
// Original definition is following
//
//	st_table *rb_st_init_strtable(void)
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/st.h
func RbStInitStrtable() *StTable {
	ret := (*StTable)(C.rb_st_init_strtable())
	return ret
}

// RbStInitStrtableWithSize calls `rb_st_init_strtable_with_size` in C
//
// Original definition is following
//
//	st_table *rb_st_init_strtable_with_size(st_index_t)
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/st.h
func RbStInitStrtableWithSize(arg1 StIndexT) *StTable {
	ret := (*StTable)(C.rb_st_init_strtable_with_size(C.st_index_t(arg1)))
	return ret
}

// RbStInitTable calls `rb_st_init_table` in C
//
// Original definition is following
//
//	st_table *rb_st_init_table(const struct st_hash_type *)
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/st.h
func RbStInitTable(arg1 *StHashType) *StTable {
	var cArg1 C.st_hash_type
	ret := *StTable(C.rb_st_init_table(&cArg1))
	*arg1 = StHashType(cArg1)
	return ret
}

// RbStInitTableWithSize calls `rb_st_init_table_with_size` in C
//
// Original definition is following
//
//	st_table *rb_st_init_table_with_size(const struct st_hash_type *, st_index_t)
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/st.h
func RbStInitTableWithSize(arg1 *StHashType, arg2 StIndexT) *StTable {
	var cArg1 C.st_hash_type
	ret := *StTable(C.rb_st_init_table_with_size(&cArg1, C.st_index_t(arg2)))
	*arg1 = StHashType(cArg1)
	return ret
}

// RbStInsert calls `rb_st_insert` in C
//
// Original definition is following
//
//	int rb_st_insert(st_table *, st_data_t, st_data_t)
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/st.h
func RbStInsert(arg1 *StTable, arg2 StDataT, arg3 StDataT) int {
	var cArg1 C.st_table
	ret := int(C.rb_st_insert(&cArg1, C.st_data_t(arg2), C.st_data_t(arg3)))
	*arg1 = StTable(cArg1)
	return ret
}

// RbStInsert2 calls `rb_st_insert2` in C
//
// Original definition is following
//
//	int rb_st_insert2(st_table *, st_data_t, st_data_t, st_data_t (*)(st_data_t))
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/st.h
func RbStInsert2(arg1 *StTable, arg2 StDataT, arg3 StDataT, arg4 unsafe.Pointer) int {
	var cArg1 C.st_table
	ret := int(C.rb_st_insert2(&cArg1, C.st_data_t(arg2), C.st_data_t(arg3), toCPointer(arg4)))
	*arg1 = StTable(cArg1)
	return ret
}

// RbStKeys calls `rb_st_keys` in C
//
// Original definition is following
//
//	st_index_t rb_st_keys(st_table *table, st_data_t *keys, st_index_t size)
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/st.h
func RbStKeys(table *StTable, keys *StDataT, size StIndexT) StIndexT {
	var cTable C.st_table
	var cKeys C.st_data_t
	ret := StIndexT(C.rb_st_keys(&cTable, &cKeys, C.st_index_t(size)))
	*table = StTable(cTable)
	*keys = StDataT(cKeys)
	return ret
}

// RbStKeysCheck calls `rb_st_keys_check` in C
//
// Original definition is following
//
//	st_index_t rb_st_keys_check(st_table *table, st_data_t *keys, st_index_t size, st_data_t never)
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/st.h
func RbStKeysCheck(table *StTable, keys *StDataT, size StIndexT, never StDataT) StIndexT {
	var cTable C.st_table
	var cKeys C.st_data_t
	ret := StIndexT(C.rb_st_keys_check(&cTable, &cKeys, C.st_index_t(size), C.st_data_t(never)))
	*table = StTable(cTable)
	*keys = StDataT(cKeys)
	return ret
}

// RbStLookup calls `rb_st_lookup` in C
//
// Original definition is following
//
//	int rb_st_lookup(st_table *, st_data_t, st_data_t *)
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/st.h
func RbStLookup(arg1 *StTable, arg2 StDataT, arg3 *StDataT) int {
	var cArg1 C.st_table
	var cArg3 C.st_data_t
	ret := int(C.rb_st_lookup(&cArg1, C.st_data_t(arg2), &cArg3))
	*arg1 = StTable(cArg1)
	*arg3 = StDataT(cArg3)
	return ret
}

// RbStReplace calls `rb_st_replace` in C
//
// Original definition is following
//
//	st_table *rb_st_replace(st_table *new_tab, st_table *old_tab)
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/st.h
func RbStReplace(new_tab *StTable, old_tab *StTable) *StTable {
	var cNewTab C.st_table
	var cOldTab C.st_table
	ret := (*StTable)(C.rb_st_replace(&cNewTab, &cOldTab))
	*new_tab = StTable(cNewTab)
	*old_tab = StTable(cOldTab)
	return ret
}

// RbStShift calls `rb_st_shift` in C
//
// Original definition is following
//
//	int rb_st_shift(st_table *, st_data_t *, st_data_t *); \/* returns 0:notfound 1:deleted *\/
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/st.h
func RbStShift(arg1 *StTable, arg2 *StDataT, arg3 *StDataT) int {
	var cArg1 C.st_table
	var cArg2 C.st_data_t
	var cArg3 C.st_data_t
	ret := int(C.rb_st_shift(&cArg1, &cArg2, &cArg3))
	*arg1 = StTable(cArg1)
	*arg2 = StDataT(cArg2)
	*arg3 = StDataT(cArg3)
	return ret
}

// RbStTableSize calls `rb_st_table_size` in C
//
// Original definition is following
//
//	size_t rb_st_table_size(const struct st_table *tbl)
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/st.h
func RbStTableSize(tbl *StTable) SizeT {
	var cTbl C.st_table
	ret := SizeT(C.rb_st_table_size(&cTbl))
	*tbl = StTable(cTbl)
	return ret
}

// RbStUpdate calls `rb_st_update` in C
//
// Original definition is following
//
//	int rb_st_update(st_table *table, st_data_t key, st_update_callback_func *func, st_data_t arg)
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/st.h
func RbStUpdate(table *StTable, key StDataT, fun *StUpdateCallbackFunc, arg StDataT) int {
	var cTable C.st_table
	var cFun C.st_update_callback_func
	ret := int(C.rb_st_update(&cTable, C.st_data_t(key), &cFun, C.st_data_t(arg)))
	*table = StTable(cTable)
	*fun = StUpdateCallbackFunc(cFun)
	return ret
}

// RbStValues calls `rb_st_values` in C
//
// Original definition is following
//
//	st_index_t rb_st_values(st_table *table, st_data_t *values, st_index_t size)
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/st.h
func RbStValues(table *StTable, values *StDataT, size StIndexT) StIndexT {
	var cTable C.st_table
	var cValues C.st_data_t
	ret := StIndexT(C.rb_st_values(&cTable, &cValues, C.st_index_t(size)))
	*table = StTable(cTable)
	*values = StDataT(cValues)
	return ret
}

// RbStValuesCheck calls `rb_st_values_check` in C
//
// Original definition is following
//
//	st_index_t rb_st_values_check(st_table *table, st_data_t *values, st_index_t size, st_data_t never)
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/st.h
func RbStValuesCheck(table *StTable, values *StDataT, size StIndexT, never StDataT) StIndexT {
	var cTable C.st_table
	var cValues C.st_data_t
	ret := StIndexT(C.rb_st_values_check(&cTable, &cValues, C.st_index_t(size), C.st_data_t(never)))
	*table = StTable(cTable)
	*values = StDataT(cValues)
	return ret
}