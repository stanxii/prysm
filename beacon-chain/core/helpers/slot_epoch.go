package helpers

import (
	pb "github.com/prysmaticlabs/prysm/proto/beacon/p2p/v1"
	"github.com/prysmaticlabs/prysm/shared/params"
)

// SlotToEpoch returns the epoch number of the input slot.
//
// Spec pseudocode definition:
//   def slot_to_epoch(slot: SlotNumber) -> EpochNumber:
//    return slot // SLOTS_PER_EPOCH
func SlotToEpoch(slot uint64) uint64 {
	return slot / params.BeaconConfig().SlotsPerEpoch
}

// CurrentEpoch returns the current epoch number calculated from
// the slot number stored in beacon state.
//
// Spec pseudocode definition:
//   def get_current_epoch(state: BeaconState) -> EpochNumber:
//    return slot_to_epoch(state.slot)
func CurrentEpoch(state *pb.BeaconState) uint64 {
	return SlotToEpoch(state.Slot)
}

// PrevEpoch returns the previous epoch number calculated from
// the slot number stored in beacon state. It also checks for
// underflow condition.
func PrevEpoch(state *pb.BeaconState) uint64 {
	if SlotToEpoch(state.Slot) == params.BeaconConfig().GenesisEpoch {
		return params.BeaconConfig().GenesisEpoch
	}
	return SlotToEpoch(state.Slot) - 1
}

// NextEpoch returns the next epoch number calculated form
// the slot number stored in beacon state.
func NextEpoch(state *pb.BeaconState) uint64 {
	return SlotToEpoch(state.Slot) + 1
}

// StartSlot returns the first slot number of the
// current epoch.
//
// Spec pseudocode definition:
//   def get_epoch_start_slot(epoch: EpochNumber) -> SlotNumber:
//    return epoch * SLOTS_PER_EPOCH
func StartSlot(epoch uint64) uint64 {
	return epoch * params.BeaconConfig().SlotsPerEpoch
}

// AttestationCurrentEpoch returns the current epoch referenced by the attestation.
func AttestationCurrentEpoch(att *pb.AttestationData) uint64 {
	return SlotToEpoch(att.Slot)
}
