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

#include "dataplane/standalone/sai/qos_map.h"

#include <glog/logging.h>

#include "dataplane/standalone/sai/common.h"
#include "dataplane/standalone/sai/entry.h"

const sai_qos_map_api_t l_qos_map = {
    .create_qos_map = l_create_qos_map,
    .remove_qos_map = l_remove_qos_map,
    .set_qos_map_attribute = l_set_qos_map_attribute,
    .get_qos_map_attribute = l_get_qos_map_attribute,
};

sai_status_t l_create_qos_map(sai_object_id_t *qos_map_id,
                              sai_object_id_t switch_id, uint32_t attr_count,
                              const sai_attribute_t *attr_list) {
  LOG(INFO) << "Func: " << __PRETTY_FUNCTION__;
  return translator->create(SAI_OBJECT_TYPE_QOS_MAP, qos_map_id, switch_id,
                            attr_count, attr_list);
}

sai_status_t l_remove_qos_map(sai_object_id_t qos_map_id) {
  LOG(INFO) << "Func: " << __PRETTY_FUNCTION__;
  return translator->remove(SAI_OBJECT_TYPE_QOS_MAP, qos_map_id);
}

sai_status_t l_set_qos_map_attribute(sai_object_id_t qos_map_id,
                                     const sai_attribute_t *attr) {
  LOG(INFO) << "Func: " << __PRETTY_FUNCTION__;
  return translator->set_attribute(SAI_OBJECT_TYPE_QOS_MAP, qos_map_id, attr);
}

sai_status_t l_get_qos_map_attribute(sai_object_id_t qos_map_id,
                                     uint32_t attr_count,
                                     sai_attribute_t *attr_list) {
  LOG(INFO) << "Func: " << __PRETTY_FUNCTION__;
  return translator->get_attribute(SAI_OBJECT_TYPE_QOS_MAP, qos_map_id,
                                   attr_count, attr_list);
}
