package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println("teste")
	// cmd := exec.Command("java", "-classpath", "/mnt/sda4/go-work/src/github.com/paulorfarah/repos/TestProject/target/classes:$RANDOOP_JAR", "randoop.main.Main", "gentests", "--testclass=testproject.Test")
	// script := CreateRandoopScript("testproject.Test")
	// cmd := exec.Command("bash " + script)
	c := "java -classpath /mnt/sda4/go-work/src/github.com/paulorfarah/repos/TestProject/target/classes:$RANDOOP_JAR randoop.main.Main gentests --testclass=testproject.Test"
	cmd := exec.Command("bash", "-c", c)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("\n[>>ERROR]: Cannot run randoop gentests (" + fmt.Sprint(err) + "): " + stderr.String())
		fmt.Println(out)
	} else {
		fmt.Println("\n [>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>SUCCESS]: Randoop executed successully!")
		fmt.Println(out.String())

	}
}

func CreateRandoopScript(class string) string {
	fn := strings.ReplaceAll(class, ".", "_") + ".sh"
	// Create new file
	f, err := os.Create(fn)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = os.Chmod(fn, 0700)
	if err != nil {
		log.Fatal(err)
	}

	c := "java -classpath /mnt/sda4/go-work/src/github.com/paulorfarah/repos/TestProject/target/classes:$RANDOOP_JAR randoop.main.Main gentests --testclass=" + class + " > " + class + ".txt"
	_, err2 := f.WriteString(c)

	if err2 != nil {
		log.Fatal(err2)
	}

	return fn
}
