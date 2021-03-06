package cpi

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/rackhd/rackhd-cpi/bosh"
	"github.com/rackhd/rackhd-cpi/config"
	"github.com/rackhd/rackhd-cpi/rackhdapi"
)

func DetachDisk(c config.Cpi, extInput bosh.MethodArguments) error {
	var vmCID string
	var diskCID string

	if reflect.TypeOf(extInput[0]) != reflect.TypeOf(vmCID) {
		return errors.New("Received unexpected type for vm cid")
	}

	if reflect.TypeOf(extInput[1]) != reflect.TypeOf(diskCID) {
		return errors.New("Received unexpected type for disk cid")
	}

	vmCID = extInput[0].(string)
	diskCID = extInput[1].(string)

	nodes, err := rackhdapi.GetNodes(c)
	if err != nil {
		return err
	}

	for _, node := range nodes {
		if node.PersistentDisk.DiskCID == diskCID {
			if !node.PersistentDisk.IsAttached {
				return fmt.Errorf("Disk: %s is detached\n", diskCID)
			}

			if node.CID != vmCID {
				return fmt.Errorf("Disk %s does not belong to VM %s\n", diskCID, vmCID)
			}

			return rackhdapi.MakeDiskRequest(c, node, false)
		}
	}

	return fmt.Errorf("Disk: %s not found\n", diskCID)
}
