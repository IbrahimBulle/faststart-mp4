package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main(){
	homedir,err:=os.UserHomeDir()
	if err!=nil{
		log.Fatal(err)
	}
    var location string
	fmt.Println("Where are your videos? (e.g. Downloads): ")
	fmt.Scan(&location)
	downloadspath:=filepath.Join(homedir,location)

	err=filepath.Walk(downloadspath,func(path string, info fs.FileInfo, err error) error {
     if !info.IsDir() && filepath.Ext(path)==".mp4"{
		fmt.Println("Processing:", path)
        tempfile:=path+".temp.mp4"
		 cmd:=exec.Command(
		"ffmpeg",
		"-i", path,
		"-c","copy",
		"-movflags","faststart",
		 tempfile,
	   )
	   cmd.Stdout = os.Stdout
       cmd.Stderr = os.Stderr

	   	if err := cmd.Run(); err != nil {
				fmt.Println("Error processing", path, ":", err)
				return err
			}

			
				os.Remove(path)           

			if err := os.Rename(tempfile, path); err != nil {
				fmt.Println("Error replacing file:", err)
				return err
			}
     }
   
   return nil
	})
	if err!=nil{
		log.Fatal(err)
	}
		fmt.Println("All MP4 files processed with fast start!")

   
}