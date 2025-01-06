// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated
// Code generated by ruby_h_to_go. DO NOT EDIT.

//go:build ruby_3_4

package ruby

/*
#include "ruby.h"
*/
import "C"

// ID is a type for passing `C.ID` in and out of package
type ID C.ID

// VALUE is a type for passing `C.VALUE` in and out of package
type VALUE C.VALUE

// OffT is a type for passing `C.off_t` in and out of package
type OffT C.off_t

// PidT is a type for passing `C.pid_t` in and out of package
type PidT C.pid_t

// RbAllocFuncT is a type for passing `C.rb_alloc_func_t` in and out of package
type RbAllocFuncT C.rb_alloc_func_t

// RbArithmeticSequenceComponentsT is a type for passing `C.rb_arithmetic_sequence_components_t` in and out of package
type RbArithmeticSequenceComponentsT C.rb_arithmetic_sequence_components_t

// RbBlockCallFunc is a type for passing `C.rb_block_call_func` in and out of package
type RbBlockCallFunc C.rb_block_call_func

// RbBlockCallFuncT is a type for passing `C.rb_block_call_func_t` in and out of package
type RbBlockCallFuncT C.rb_block_call_func_t

// RbBlockingFunctionT is a type for passing `C.rb_blocking_function_t` in and out of package
type RbBlockingFunctionT C.rb_blocking_function_t

// RbDataTypeT is a type for passing `C.rb_data_type_t` in and out of package
type RbDataTypeT C.rb_data_type_t

// RbEnumeratorSizeFunc is a type for passing `C.rb_enumerator_size_func` in and out of package
type RbEnumeratorSizeFunc C.rb_enumerator_size_func

// RbEventFlagT is a type for passing `C.rb_event_flag_t` in and out of package
type RbEventFlagT C.rb_event_flag_t

// RbEventHookFuncT is a type for passing `C.rb_event_hook_func_t` in and out of package
type RbEventHookFuncT C.rb_event_hook_func_t

// RbFdsetT is a type for passing `C.rb_fdset_t` in and out of package
type RbFdsetT C.rb_fdset_t

// RbGvarGetterT is a type for passing `C.rb_gvar_getter_t` in and out of package
type RbGvarGetterT C.rb_gvar_getter_t

// RbGvarMarkerT is a type for passing `C.rb_gvar_marker_t` in and out of package
type RbGvarMarkerT C.rb_gvar_marker_t

// RbGvarSetterT is a type for passing `C.rb_gvar_setter_t` in and out of package
type RbGvarSetterT C.rb_gvar_setter_t

// RbHashUpdateFunc is a type for passing `C.rb_hash_update_func` in and out of package
type RbHashUpdateFunc C.rb_hash_update_func

// RbUnblockFunctionT is a type for passing `C.rb_unblock_function_t` in and out of package
type RbUnblockFunctionT C.rb_unblock_function_t

// RbWarningCategoryT is a type for passing `C.rb_warning_category_t` in and out of package
type RbWarningCategoryT C.rb_warning_category_t

// StCheckForSizeofStIndexT is a type for passing `C.st_check_for_sizeof_st_index_t` in and out of package
type StCheckForSizeofStIndexT C.st_check_for_sizeof_st_index_t

// StCompareFunc is a type for passing `C.st_compare_func` in and out of package
type StCompareFunc C.st_compare_func

// StDataT is a type for passing `C.st_data_t` in and out of package
type StDataT C.st_data_t

// StForeachCallbackFunc is a type for passing `C.st_foreach_callback_func` in and out of package
type StForeachCallbackFunc C.st_foreach_callback_func

// StForeachCheckCallbackFunc is a type for passing `C.st_foreach_check_callback_func` in and out of package
type StForeachCheckCallbackFunc C.st_foreach_check_callback_func

// StHashFunc is a type for passing `C.st_hash_func` in and out of package
type StHashFunc C.st_hash_func

// StIndexT is a type for passing `C.st_index_t` in and out of package
type StIndexT C.st_index_t

// StTable is a type for passing `C.st_table` in and out of package
type StTable C.st_table

// StTableEntry is a type for passing `C.st_table_entry` in and out of package
type StTableEntry C.st_table_entry

// StUpdateCallbackFunc is a type for passing `C.st_update_callback_func` in and out of package
type StUpdateCallbackFunc C.st_update_callback_func
