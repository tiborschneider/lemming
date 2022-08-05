#!/bin/bash
#
# Copyright 2021 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# This script is used to generate the Ondatra Telemetry and Config Go APIs.

set -e

go install github.com/openconfig/ygot/generator@latest
git clone https://github.com/openconfig/public.git

EXCLUDE_MODULES=ietf-interfaces,openconfig-bfd,openconfig-messages

COMMON_ARGS=(
  -path="public/release/models,public/third_party/ietf"
  -compress_paths
  -exclude_modules="${EXCLUDE_MODULES}"
  -generate_fakeroot
  -fakeroot_name=device
  -generate_simple_unions
  -shorten_enum_leaf_names
  -typedef_enum_with_defmod
  -enum_suffix_for_simple_union_enums
  -trim_enum_openconfig_prefix
  -include_schema
  -generate_append
  -generate_getters
  -generate_rename
  -generate_delete
  -generate_leaf_getters
  -structs_split_files_count=10
  -generate_populate_defaults
)

YANG_FILES=(
  public/release/models/acl/openconfig-acl.yang
  public/release/models/acl/openconfig-packet-match.yang
  public/release/models/aft/openconfig-aft.yang
  public/release/models/bfd/openconfig-bfd.yang
  public/release/models/bgp/openconfig-bgp-policy.yang
  public/release/models/bgp/openconfig-bgp-types.yang
  public/release/models/interfaces/openconfig-if-aggregate.yang
  public/release/models/interfaces/openconfig-if-ethernet.yang
  public/release/models/interfaces/openconfig-if-ip-ext.yang
  public/release/models/interfaces/openconfig-if-ip.yang
  public/release/models/interfaces/openconfig-interfaces.yang
  public/release/models/isis/openconfig-isis.yang
  public/release/models/lacp/openconfig-lacp.yang
  public/release/models/lldp/openconfig-lldp-types.yang
  public/release/models/lldp/openconfig-lldp.yang
  public/release/models/local-routing/openconfig-local-routing.yang
  public/release/models/mpls/openconfig-mpls-types.yang
  public/release/models/multicast/openconfig-pim.yang
  public/release/models/network-instance/openconfig-network-instance.yang
  public/release/models/openconfig-extensions.yang
  public/release/models/optical-transport/openconfig-transport-types.yang
  public/release/models/ospf/openconfig-ospfv2.yang
  public/release/models/platform/openconfig-platform-cpu.yang
  public/release/models/platform/openconfig-platform-integrated-circuit.yang
  public/release/models/platform/openconfig-platform-software.yang
  public/release/models/platform/openconfig-platform-transceiver.yang
  public/release/models/platform/openconfig-platform.yang
  public/release/models/policy-forwarding/openconfig-policy-forwarding.yang
  public/release/models/policy/openconfig-policy-types.yang
  public/release/models/qos/openconfig-qos-elements.yang
  public/release/models/qos/openconfig-qos-interfaces.yang
  public/release/models/qos/openconfig-qos-types.yang
  public/release/models/qos/openconfig-qos.yang
  public/release/models/rib/openconfig-rib-bgp.yang
  public/release/models/segment-routing/openconfig-segment-routing-types.yang
  public/release/models/system/openconfig-system.yang
  public/release/models/types/openconfig-inet-types.yang
  public/release/models/types/openconfig-types.yang
  public/release/models/types/openconfig-yang-types.yang
  public/release/models/vlan/openconfig-vlan.yang
  public/third_party/ietf/iana-if-type.yang
  public/third_party/ietf/ietf-inet-types.yang
  public/third_party/ietf/ietf-interfaces.yang
  public/third_party/ietf/ietf-yang-types.yang
)

# Generate Config Structs and Path API
mkdir -p config/device
echo "generating structs for config"
generator \
  -generate_structs \
  -generate_path_structs=false \
  -list_builder_key_threshold=4 \
  -output_dir=config \
  -package_name=config \
  "${COMMON_ARGS[@]}" \
  "${YANG_FILES[@]}"

echo "generating path structs for config"
generator \
  -generate_structs=false \
  -generate_path_structs=true \
  -exclude_state \
  -schema_struct_path=github.com/openconfig/lemming/gnmi/internal/config \
  -output_dir=config \
  -package_name=device \
  -split_pathstructs_by_module=true \
  -path_structs_output_file=config/device/device.go \
  -base_import_path=github.com/openconfig/lemming/gnmi/internal/config \
  -trim_path_package_prefix="openconfig" \
  -path_struct_package_suffix="" \
  "${COMMON_ARGS[@]}" \
  "${YANG_FILES[@]}"

# Generate State Structs and Path API
mkdir -p telemetry/device
echo "generating structs for telemetry"
generator \
  -generate_structs \
  -generate_path_structs=false \
  -prefer_operational_state \
  -list_builder_key_threshold=4 \
  -output_dir=telemetry \
  -package_name=telemetry \
  "${COMMON_ARGS[@]}" \
  "${YANG_FILES[@]}"

echo "generating path structs for telemetry"
generator \
  -generate_structs=false \
  -generate_path_structs=true \
  -prefer_operational_state \
  -list_builder_key_threshold=4 \
  -output_dir=telemetry \
  -package_name=device \
  -path_structs_output_file=telemetry/device/device.go \
  -split_pathstructs_by_module=true \
  -schema_struct_path=github.com/openconfig/lemming/gnmi/internal/telemetry \
  -trim_path_package_prefix="openconfig" \
  -path_struct_package_suffix="" \
  -base_import_path=github.com/openconfig/lemming/gnmi/internal/telemetry \
  "${COMMON_ARGS[@]}" \
  "${YANG_FILES[@]}"

find config telemetry -name "*.go" -exec goimports -w {} +
find config telemetry -name "*.go" -exec gofmt -w -s {} +
rm -rf public
