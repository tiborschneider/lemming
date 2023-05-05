
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

#ifndef DATAPLANE_STANDALONE_SAI_MPLS_H_
#define DATAPLANE_STANDALONE_SAI_MPLS_H_

extern "C" {
#include "inc/sai.h"
}

extern const sai_mpls_api_t l_mpls;

sai_status_t l_create_inseg_entry(const sai_inseg_entry_t *inseg_entry,
                                  uint32_t attr_count,
                                  const sai_attribute_t *attr_list);

sai_status_t l_remove_inseg_entry(const sai_inseg_entry_t *inseg_entry);

sai_status_t l_set_inseg_entry_attribute(const sai_inseg_entry_t *inseg_entry,
                                         const sai_attribute_t *attr);

sai_status_t l_get_inseg_entry_attribute(const sai_inseg_entry_t *inseg_entry,
                                         uint32_t attr_count,
                                         sai_attribute_t *attr_list);

#endif  // DATAPLANE_STANDALONE_SAI_MPLS_H_
