
// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

#include "dataplane/standalone/sai/lag.h"

#include "dataplane/standalone/log/log.h"

const sai_lag_api_t l_lag = {
    .create_lag = l_create_lag,
    .remove_lag = l_remove_lag,
    .set_lag_attribute = l_set_lag_attribute,
    .get_lag_attribute = l_get_lag_attribute,
    .create_lag_member = l_create_lag_member,
    .remove_lag_member = l_remove_lag_member,
    .set_lag_member_attribute = l_set_lag_member_attribute,
    .get_lag_member_attribute = l_get_lag_member_attribute,
    .create_lag_members = l_create_lag_members,
    .remove_lag_members = l_remove_lag_members,
};

sai_status_t l_create_lag(sai_object_id_t *lag_id, sai_object_id_t switch_id,
                          uint32_t attr_count,
                          const sai_attribute_t *attr_list) {
  LUCIUS_LOG_FUNC();
  return SAI_STATUS_NOT_IMPLEMENTED;
}

sai_status_t l_remove_lag(sai_object_id_t lag_id) {
  LUCIUS_LOG_FUNC();
  return SAI_STATUS_NOT_IMPLEMENTED;
}

sai_status_t l_set_lag_attribute(sai_object_id_t lag_id,
                                 const sai_attribute_t *attr) {
  LUCIUS_LOG_FUNC();
  return SAI_STATUS_NOT_IMPLEMENTED;
}

sai_status_t l_get_lag_attribute(sai_object_id_t lag_id, uint32_t attr_count,
                                 sai_attribute_t *attr_list) {
  LUCIUS_LOG_FUNC();
  return SAI_STATUS_NOT_IMPLEMENTED;
}

sai_status_t l_create_lag_member(sai_object_id_t *lag_member_id,
                                 sai_object_id_t switch_id, uint32_t attr_count,
                                 const sai_attribute_t *attr_list) {
  LUCIUS_LOG_FUNC();
  return SAI_STATUS_NOT_IMPLEMENTED;
}

sai_status_t l_remove_lag_member(sai_object_id_t lag_member_id) {
  LUCIUS_LOG_FUNC();
  return SAI_STATUS_NOT_IMPLEMENTED;
}

sai_status_t l_set_lag_member_attribute(sai_object_id_t lag_member_id,
                                        const sai_attribute_t *attr) {
  LUCIUS_LOG_FUNC();
  return SAI_STATUS_NOT_IMPLEMENTED;
}

sai_status_t l_get_lag_member_attribute(sai_object_id_t lag_member_id,
                                        uint32_t attr_count,
                                        sai_attribute_t *attr_list) {
  LUCIUS_LOG_FUNC();
  return SAI_STATUS_NOT_IMPLEMENTED;
}

sai_status_t l_create_lag_members(sai_object_id_t switch_id,
                                  uint32_t object_count,
                                  const uint32_t *attr_count,
                                  const sai_attribute_t **attr_list,
                                  sai_bulk_op_error_mode_t mode,
                                  sai_object_id_t *object_id,
                                  sai_status_t *object_statuses) {
  LUCIUS_LOG_FUNC();
  return SAI_STATUS_NOT_IMPLEMENTED;
}

sai_status_t l_remove_lag_members(uint32_t object_count,
                                  const sai_object_id_t *object_id,
                                  sai_bulk_op_error_mode_t mode,
                                  sai_status_t *object_statuses) {
  LUCIUS_LOG_FUNC();
  return SAI_STATUS_NOT_IMPLEMENTED;
}
