package main

import (
    "fmt"
    "strings"
    "os"
    "os/exec"
    "cpu"
)

func getLogo() {
}
func getDistro() string{
	distro, err := os.Open("/etc/os-release")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(0)
	}
	distro_info := make([]byte, 1024)
	distro.Read(distro_info)
	distro.Close()
	distro_list := strings.Split(string(distro_info), "\n")
	distro_tuples := make(map[string]string)
	for _, v := range distro_list {
		if strings.Contains(v, "=") {
			kv := strings.Split(v, "=")
			kv[0] = strings.TrimSpace(kv[0])
			kv[1] = strings.TrimSpace(kv[1])
			distro_tuples[kv[0]] = kv[1]
		}
	}
	return strings.Trim(distro_tuples["PRETTY_NAME"], "\"") 
}
func getHost() {
}
func getKernel() string {
    cmd := exec.Command("uname","-r")
    kernel, err := cmd.Output()
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    return string(kernel)
}
func getUptime() {
}
func getPackages() {
}
func getShell() string {
    cmd := exec.Command("echo","$SHELL")
    shell, err := cmd.Output()
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    return string(shell)

}
func getResolution() {
}
func getWM() {
}
func getTheme() {
}
func getIcons() {
}
func getTerminal() {
}
func getCPU() {
}
func getGPU() {
}
func getMemory() {
      
      //
      // The coolest part about this function unlike neofetch is that it also takes account of the basic os operations
      // into account thus not giving an illusion that your os is completely having fun
      // honestly idk lmao
      //
      
      mem, err := os.Open("/proc/meminfo")
      if err != nil {
          fmt.Println(err.Error())
          os.Exit(0)
      }
      mem_info := make([]byte, 1024)
      mem.Read(mem_info)
      mem.Close()
      mem_list := strings.Split(string(mem_info),"\n");
      mem_map := make(map[string]string)
      for _,v := range mem_list {
          if strings.Contains(v,":") {
              kv := strings.Split(v,":")
              kv[0] = strings.TrimSpace(kv[0])
              kv[1] = strings.TrimSpace(kv[1])
              kv[1] = strings.Replace(kv[1],"kB","",3)
              kv[1] = strings.TrimSpace(kv[1])
              mem_map[kv[0]] = kv[1]
          }
      }
      mem_free,_ :=  strconv.Atoi(mem_map["MemFree"])
      mem_total,_ := strconv.Atoi(mem_map["MemTotal"])
      mem_used := mem_total-mem_free                                                                                                                                                                          
      memory := fmt.Sprintf("%d/%d", mem_used/1024, mem_total/1024)      
      return memory 
}

func getColorPalette() {
}


func main() {


}
