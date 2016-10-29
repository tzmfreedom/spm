package main

import (
  "os"
  "path/filepath"
  "fmt"
  "io"

  "github.com/urfave/cli"
  "gopkg.in/src-d/go-git.v4"
)

type Config struct {
  username string
  password string
  endpoint string
  api_version string
}

func main() {
  app := cli.NewApp()
  app.Name = "spm"
  app.Usage = "salesforce package manager"
  app.Commands = []cli.Command{
    {
      Name:    "install",
      Aliases: []string{"i"},
      Usage:   "install salesforce package",
      //Flags:   []cli.Flag {
      //  cli.StringFlag{
      //    Name: "lang, l",
      //    Value: "english",
      //    Usage: "language for the greeting",
      //  },
      //},
      Action:  func(c *cli.Context) error {
        var username string
        var password string
        fmt.Print("Enter Username: ")
        fmt.Scanln(&username)
        fmt.Printf("Enter Password: ")
        fmt.Scanln(&password)
        fmt.Printf("%q %q", username, password)
        os.Exit(0)

        install("https://" + c.Args().First(), nil)
        return nil
      },
    },
  }

  app.Run(os.Args)
}

func install(url string, options []string) {
  cloneFromRemoteRepository(url, url, options)
  deployToSalesforce(url, Config{})
}

func cloneFromRemoteRepository(directory string, url string, option []string) {
  r, err := git.NewFilesystemRepository(directory)
  err = r.Clone(&git.CloneOptions{
    URL: url,
  })
  if err != nil {
    panic(err)
  }

  // ... retrieving the branch being pointed by HEAD
  ref, _ := r.Head()

  // ... retrieving the commit object
  commit, _ := r.Commit(ref.Hash())

  // ... we get all the files from the commit
  files, _ := commit.Files()

  _, err = r.Head()
  err = files.ForEach(func(f *git.File) error {
    abs := filepath.Join(directory, f.Name)
    dir := filepath.Dir(abs)

    os.MkdirAll(dir, 0777)
    file, err := os.Create(abs)
    if err != nil {
      return err
    }

    defer file.Close()
    r, err := f.Reader()
    if err != nil {
      return err
    }

    defer r.Close()

    if err := file.Chmod(f.Mode); err != nil {
      return err
    }

    _, err = io.Copy(file, r)
    return err
  })
}

func deployToSalesforce(directory string, config Config) {

}