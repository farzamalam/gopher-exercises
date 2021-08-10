package main

type Track struct {
	Name      string
	AlbumName string
	Artist    string
}

var script = `apiVersion: v1
kind: PersistentVolume
metadata:
  name: pv00001
spec:
  capacity:
    storage: 10Gi
  accessModes:
    - ReadWriteOnce
  persistentVolumeReclaimPolicy: Retain   
  nfs:
    path: /mnt/nfs_shares/k8s/root/pv00001
    server: 192.168.10.223`

var tracks = []Track{
	{"foo", "album1", "artist1"},
	{"bar", "album1", "artist1"},
	{"pv", script, "artist1"},
	{"baz", "album2", "artist2"},
	{"baz", "album3", "artist2"},
}

// func main() {
// 	idx, err := fuzzyfinder.Find(
// 		tracks,
// 		func(i int) string {
// 			return tracks[i].Name
// 		},
// 		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
// 			if i == -1 {
// 				return ""
// 			}
// 			return fmt.Sprintf("Track: %s (%s)\nAlbum: %s",
// 				tracks[i].Name,
// 				tracks[i].Artist,
// 				tracks[i].AlbumName)
// 		}))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("selected: %v\n", tracks[idx].AlbumName)
// }
