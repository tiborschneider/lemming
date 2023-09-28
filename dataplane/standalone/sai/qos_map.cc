

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

#include "dataplane/standalone/proto/common.pb.h"
#include "dataplane/standalone/proto/qos_map.pb.h"
#include "dataplane/standalone/sai/common.h"

const sai_qos_map_api_t l_qos_map = {
    .create_qos_map = l_create_qos_map,
    .remove_qos_map = l_remove_qos_map,
    .set_qos_map_attribute = l_set_qos_map_attribute,
    .get_qos_map_attribute = l_get_qos_map_attribute,
};

lemming::dataplane::sai::CreateQosMapRequest convert_create_qos_map(
    sai_object_id_t switch_id, uint32_t attr_count,
    const sai_attribute_t *attr_list) {
  lemming::dataplane::sai::CreateQosMapRequest msg;

  for (uint32_t i = 0; i < attr_count; i++) {
    switch (attr_list[i].id) {
      case SAI_QOS_MAP_ATTR_TYPE:
        msg.set_type(static_cast<lemming::dataplane::sai::QosMapType>(
            attr_list[i].value.s32 + 1));
        break;
    }
  }
  return msg;
}

sai_status_t l_create_qos_map(sai_object_id_t *qos_map_id,
                              sai_object_id_t switch_id, uint32_t attr_count,
                              const sai_attribute_t *attr_list) {
  LOG(INFO) << "Func: " << __PRETTY_FUNCTION__;

  lemming::dataplane::sai::CreateQosMapRequest req =
      convert_create_qos_map(switch_id, attr_count, attr_list);
  lemming::dataplane::sai::CreateQosMapResponse resp;
  grpc::ClientContext context;
  req.set_switch_(switch_id);

  grpc::Status status = qos_map->CreateQosMap(&context, req, &resp);
  if (!status.ok()) {
    LOG(ERROR) << status.error_message();
    return SAI_STATUS_FAILURE;
  }
  *qos_map_id = resp.oid();

  return SAI_STATUS_SUCCESS;
}

sai_status_t l_remove_qos_map(sai_object_id_t qos_map_id) {
  LOG(INFO) << "Func: " << __PRETTY_FUNCTION__;

  lemming::dataplane::sai::RemoveQosMapRequest req;
  lemming::dataplane::sai::RemoveQosMapResponse resp;
  grpc::ClientContext context;
  req.set_oid(qos_map_id);

  grpc::Status status = qos_map->RemoveQosMap(&context, req, &resp);
  if (!status.ok()) {
    LOG(ERROR) << status.error_message();
    return SAI_STATUS_FAILURE;
  }

  return SAI_STATUS_SUCCESS;
}

sai_status_t l_set_qos_map_attribute(sai_object_id_t qos_map_id,
                                     const sai_attribute_t *attr) {
  LOG(INFO) << "Func: " << __PRETTY_FUNCTION__;

  return SAI_STATUS_SUCCESS;
}

sai_status_t l_get_qos_map_attribute(sai_object_id_t qos_map_id,
                                     uint32_t attr_count,
                                     sai_attribute_t *attr_list) {
  LOG(INFO) << "Func: " << __PRETTY_FUNCTION__;

  lemming::dataplane::sai::GetQosMapAttributeRequest req;
  lemming::dataplane::sai::GetQosMapAttributeResponse resp;
  grpc::ClientContext context;

  req.set_oid(qos_map_id);

  for (uint32_t i = 0; i < attr_count; i++) {
    req.add_attr_type(
        static_cast<lemming::dataplane::sai::QosMapAttr>(attr_list[i].id + 1));
  }
  grpc::Status status = qos_map->GetQosMapAttribute(&context, req, &resp);
  if (!status.ok()) {
    LOG(ERROR) << status.error_message();
    return SAI_STATUS_FAILURE;
  }
  for (uint32_t i = 0; i < attr_count; i++) {
    switch (attr_list[i].id) {
      case SAI_QOS_MAP_ATTR_TYPE:
        attr_list[i].value.s32 = static_cast<int>(resp.attr().type() - 1);
        break;
    }
  }

  return SAI_STATUS_SUCCESS;
}
