// Code generated by "libovsdb.modelgen"
// DO NOT EDIT.

package ovnsb

const LogicalDPGroupTable = "Logical_DP_Group"

// LogicalDPGroup defines an object in Logical_DP_Group table
type LogicalDPGroup struct {
	UUID      string   `ovsdb:"_uuid"`
	Datapaths []string `ovsdb:"datapaths"`
}
