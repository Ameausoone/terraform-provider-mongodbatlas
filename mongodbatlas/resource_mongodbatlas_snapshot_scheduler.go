package mongodbatlas

//
//import (
//	ma "github.com/akshaykarle/go-mongodbatlas/mongodbatlas"
//	"github.com/hashicorp/terraform/helper/schema"
//	"github.com/hashicorp/terraform/helper/validation"
//)
//
//// api documentation : https://docs.atlas.mongodb.com/reference/api/snapshot-schedule-patch/
//func resourceSnapshotScheduler() *schema.Resource {
//	return &schema.Resource{
//		Create: ResourceSnapshotSchedulerCreate,
//		Read:   ResourceSnapshotSchedulerRead,
//		Schema: map[string]*schema.Schema{
//			"group_id": {
//				Type:     schema.TypeString,
//				Required: true,
//				ForceNew: true,
//			},
//			"cluster_name": {
//				Type:     schema.TypeString,
//				Required: true,
//				ForceNew: true,
//			},
//
//			"snapshot_interval_hours": {
//				Type:         schema.TypeInt,
//				Required:     true,
//				ValidateFunc: validation.IntInSlice([]int{6, 8, 12, 24}),
//			},
//			"snapshot_retention_days": {
//				Type:         schema.TypeInt,
//				Required:     true,
//				ValidateFunc: validation.IntInSlice([]int{2, 3, 4, 5}),
//			},
//			"daily_snapshot_retention_days": {
//				Type:         schema.TypeInt,
//				Required:     true,
//				ValidateFunc: validation.IntInSlice([]int{0, 3, 4, 5, 6, 7, 15, 30, 60, 90, 120, 180, 360}),
//			},
//			"point_in_time_window_hours": {
//				Type:     schema.TypeInt,
//				Required: true,
//			},
//			"cluster_checkpoint_interval_min": {
//				Type:         schema.TypeInt,
//				Required:     true,
//				ValidateFunc: validation.IntInSlice([]int{15, 30, 60}),
//			},
//
//			"weekly_snapshot_retention_weeks": {
//				Type:     schema.TypeInt,
//				Required: true,
//				ValidateFunc: validation.Any(
//					validation.IntBetween(0, 8),
//					validation.IntInSlice([]int{12, 16, 20, 24, 52}),
//				),
//			},
//			"monthly_snapshot_retention_months": {
//				Type:     schema.TypeInt,
//				Required: true,
//				ValidateFunc: validation.Any(
//					validation.IntBetween(0, 13),
//					validation.IntInSlice([]int{18, 24, 36}),
//				),
//			},
//		},
//	}
//}
//
//func ResourceSnapshotSchedulerCreate(d *schema.ResourceData, meta interface{}) error {
//	client := meta.(*ma.Client)
//	d.Get()
//	client.SnapshotSchedule.Update()
//}
//
//func ResourceSnapshotSchedulerRead(d *schema.ResourceData, meta interface{}) error {
//	client := meta.(*ma.Client)
//	client.SnapshotSchedule.Get(d.Get)
//	return nil
//}
