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

#ifndef DATAPLANE_STANDALONE_SAI_UDF_H_
#define DATAPLANE_STANDALONE_SAI_UDF_H_

extern "C" {
#include "inc/sai.h"
}

extern const sai_udf_api_t l_udf;

sai_status_t l_create_udf(sai_object_id_t *udf_id, sai_object_id_t switch_id,
                          uint32_t attr_count,
                          const sai_attribute_t *attr_list);

sai_status_t l_remove_udf(sai_object_id_t udf_id);

sai_status_t l_set_udf_attribute(sai_object_id_t udf_id,
                                 const sai_attribute_t *attr);

sai_status_t l_get_udf_attribute(sai_object_id_t udf_id, uint32_t attr_count,
                                 sai_attribute_t *attr_list);

sai_status_t l_create_udf_match(sai_object_id_t *udf_match_id,
                                sai_object_id_t switch_id, uint32_t attr_count,
                                const sai_attribute_t *attr_list);

sai_status_t l_remove_udf_match(sai_object_id_t udf_match_id);

sai_status_t l_set_udf_match_attribute(sai_object_id_t udf_match_id,
                                       const sai_attribute_t *attr);

sai_status_t l_get_udf_match_attribute(sai_object_id_t udf_match_id,
                                       uint32_t attr_count,
                                       sai_attribute_t *attr_list);

sai_status_t l_create_udf_group(sai_object_id_t *udf_group_id,
                                sai_object_id_t switch_id, uint32_t attr_count,
                                const sai_attribute_t *attr_list);

sai_status_t l_remove_udf_group(sai_object_id_t udf_group_id);

sai_status_t l_set_udf_group_attribute(sai_object_id_t udf_group_id,
                                       const sai_attribute_t *attr);

sai_status_t l_get_udf_group_attribute(sai_object_id_t udf_group_id,
                                       uint32_t attr_count,
                                       sai_attribute_t *attr_list);

#endif  // DATAPLANE_STANDALONE_SAI_UDF_H_
