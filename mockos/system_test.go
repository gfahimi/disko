package mockos_test

import (
	"testing"

	"github.com/anuvu/disko"
	"github.com/anuvu/disko/mockos"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSystem(t *testing.T) {
	Convey("testing System Model", t, func() {
		sys := mockos.System("testdata/model_sys.json")
		So(sys, ShouldNotBeNil)

		Convey("Calling ScanAllDisks with no filter function should return all thh disks", func() {
			diskSet, err := sys.ScanAllDisks(func(d disko.Disk) bool {return true})
			So(err, ShouldBeNil)
			So(len(diskSet), ShouldEqual, 6)

			// Todo Ravchama and gfahimi. ScanAllDisk does not in any case return error.
			// Probably it should return error when there is no found disks to be compatible
			// with the rest of functionality
		})

		Convey("Calling ScanDisk on dev/sda path should return the disk(s) with similar path ", func() {
			disk, err := sys.ScanDisk( "/dev/sda")
			So(err, ShouldBeNil)
			So(disk, ShouldNotBeNil)
			So(disk.Name, ShouldEqual, "sda")
		})

		Convey("Calling ScanDisk on path that does not contain any disk should return error ", func() {
			_, err := sys.ScanDisk( "path/with/no/disk")
			So(err, ShouldNotBeNil)
		})

		Convey("Calling ScanDisk on dev/sda path should return the disk(s) with similar path", func() {
			disk, err := sys.ScanDisk( "/dev/sda")
			So(err, ShouldBeNil)
			So(disk, ShouldNotBeNil)
			So(disk.Name, ShouldEqual, "sda")
		})

		Convey("Calling ScanDisks with a specific filter on dev/sda path should return corresponding disks", func() {
			disk, err := sys.ScanDisks(func(d disko.Disk) bool {return d.Size > 10000}, "/dev/sda")
			So(err, ShouldBeNil)
			So(disk, ShouldNotBeNil)
		})

		Convey("Calling ScanDisks with a specific filter on invalid path should return error", func() {
			disk, err := sys.ScanDisks(func(d disko.Disk) bool {return d.Size > 10000},
			"/dev/sda", "path/with/no/disk")
			So(err, ShouldNotBeNil)
			So(disk, ShouldBeNil)
		})

		Convey("Calling CreatePartition should create a partition in disk", func() {
			disk := disko.Disk{
				Name: "sda",
			}
			partition := disko.Partition{
				Start: 0,
				End: 10000,
				ID: "1234567",
				Name: "partition1",
				Number: 1,
			}
			// TODO: CreatePartition should probably only get the name
			err := sys.CreatePartition(disk, partition)
			So(err, ShouldBeNil)

			d, _ := sys.ScanDisk("/dev/sda")
			So(len(d.Partitions), ShouldEqual, 1)
			_, ok := d.Partitions[1]
			So(ok, ShouldBeTrue)

			Convey("Calling DeletePartition should delete the partition with the specific number from a disk", func() {
				disk := disko.Disk{
					Name: "sda",
				}

				// TODO: DeletePartition should probably only get the name
				err := sys.DeletePartition(disk, 1)
				So(err, ShouldBeNil)

				d, _ := sys.ScanDisk("/dev/sda")
				So(len(d.Partitions), ShouldEqual, 0)
			})

			Convey("Calling CreatePartition with an existing partition should return error", func() {
				err := sys.CreatePartition(disk, partition)
				So(err, ShouldNotBeNil)
			})
		})

		Convey("Calling CreatePartition on a disk not being track by system should return error", func() {
			disk := disko.Disk{
				Name: "invalid",
			}
			partition := disko.Partition{
				Start:  0,
				End:    10000,
				ID:     "1234567",
				Name:   "partition1",
				Number: 1,
			}
			err := sys.CreatePartition(disk, partition)
			So(err, ShouldNotBeNil)
		})


	})


}
