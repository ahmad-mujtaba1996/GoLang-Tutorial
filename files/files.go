package main

func main() {
	//=============================================Read file info
	// f, err := os.Open("example.txt") // In go, if we get an error in function we return it rather than throwing an exception
	// if err != nil {
	// 	// log the error
	// 	panic(err)
	// }

	// fileinfo, err := f.Stat() //Stat returns the [FileInfo] structure describing file. If there is an error, it will be of type [*PathError].
	// if err != nil {
	// 	// log the error
	// 	panic(err)
	// }

	// fmt.Println("file name:", fileinfo.Name())
	// fmt.Println("file folder or directory:", fileinfo.IsDir())
	// fmt.Println("file size:", fileinfo.Size())
	// fmt.Println("file permissions:", fileinfo.Mode())
	// fmt.Println("file modification time:", fileinfo.ModTime())

	//=============================================Read files
	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	// log the error
	// 	panic(err)
	// }

	// defer f.Close()

	// buff := make([]byte, 12) //Storing read values in buffer array. Right now its static but can be done dynamically based on file size like fileInfo.Size()

	// fileRead, err := f.Read(buff)
	// if err != nil {
	// 	// log the error
	// 	panic(err)
	// }

	// // fmt.Println("Buffer Data:", fileRead, string(buff)) // using string(buff) to convert byte array to string

	// for i := 0; i < len(buff); i++ {
	// 	fmt.Println("Byte:", fileRead, i, string(buff[i]))
	// }

	// file, error := os.ReadFile("example.txt") // Can read file using this approach but it is viable only with small files not with bigger files as this operation loads all file size in memory which can reduce application resources hence impacting performance
	// if error != nil {
	// 	panic(error)
	// }
	// fmt.Println("File Content:", string(file))

	// =======================================Read Folders
	// dir, error := os.Open("..") // . represents current directory. ../ means one level up in directory
	// if error != nil {
	// 	panic(error)
	// }

	// defer dir.Close()

	// fileInfo, error := dir.ReadDir(0) // n represents how many file in that directory we want to read. If n <= 0 then it will read all files in that directory
	// if error != nil {
	// 	panic(error)
	// }

	// for _, file := range fileInfo {
	// 	fmt.Println("File Name:", file.Name(), file.IsDir())
	// }

	//=======================================Create a file
	// file, error := os.Create("example2.txt")
	// if error != nil {
	// 	panic(error)
	// }
	// defer file.Close()

	//====================================== Writing String to File
	// file.WriteString("Hi GoLang!, This is Ahmad Mujtaba\n")
	// file.WriteString("Here is my miserable life")

	//======================================== Writing Binary Data to file
	// bytes := []byte("Hello Golang!")
	// file.Write(bytes)

	// Using Streaming fashion, transfer data from one file to another file.
	// sourceFile, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer sourceFile.Close()

	// destinationFile, error := os.Create("example2.txt")
	// if error != nil {
	// 	panic(error)
	// }
	// defer destinationFile.Close()

	// bufio package provides a predefined buffer that can be used to read data from file
	// reader := bufio.NewReader(sourceFile)
	// writer := bufio.NewWriter(destinationFile)

	// for {
	// 	readByte, err := reader.ReadByte()
	// 	if err != nil {
	// 		if err.Error() != "EOF" {
	// 			panic(err)
	// 		}
	// 		break
	// 	}
	// 	error := writer.WriteByte(readByte)
	// 	if error != nil {
	// 		panic(error)
	// 	}
	// }

	// // Flush the buffer to the destination file
	// e := writer.Flush()
	// if e != nil {
	// 	panic(e)
	// }

	// fmt.Println("File copied successfully")

	//====================================== Copy file content from source to destination
	// _, err = io.Copy(destinationFile, sourceFile)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("File copied successfully")

	// =====================================Delete file
	// error := os.Remove("example2.txt")
	// if error != nil {
	// 	panic(error)
	// }
	// fmt.Println("File deleted successfully")
}
