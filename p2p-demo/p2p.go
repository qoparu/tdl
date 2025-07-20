package main

import (
    "bufio"
    "fmt"
    "io"
    "log"
    "net"
    "os"
    "strings"
    "sync"
)

// Храним все соединения с пирами
var peers = make(map[string]net.Conn)
var mu sync.Mutex

// Приём сообщений от одного peer
func handleConn(conn net.Conn) {
    defer conn.Close()
    reader := bufio.NewReader(conn)
    for {
        msg, err := reader.ReadString('\n')
        if err != nil {
            if err != io.EOF {
                log.Println("Ошибка чтения:", err)
            }
            mu.Lock()
            delete(peers, conn.RemoteAddr().String())
            mu.Unlock()
            return
        }
        fmt.Print("[P2P] ", msg)
    }
}

// Подключение к другому peer
func connectToPeer(addr string) {
    conn, err := net.Dial("tcp", addr)
    if err != nil {
        log.Println("Не удалось подключиться к peer:", err)
        return
    }
    mu.Lock()
    peers[conn.RemoteAddr().String()] = conn
    mu.Unlock()
    go handleConn(conn)
}

func main() {
    // 1. Введи порт для прослушивания (например, 7001)
    fmt.Print("Введите порт для прослушивания (например, 7001): ")
    var port string
    fmt.Scanln(&port)

    // 2. Запустить listener
    ln, err := net.Listen("tcp", ":"+port)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Слушаю на порту", port)

    // 3. Можно подключиться к другому peer (опционально)
    fmt.Print("IP:порт peer для подключения (например, 127.0.0.1:7002), Enter чтобы пропустить: ")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    addr := strings.TrimSpace(scanner.Text())
    if addr != "" {
        go connectToPeer(addr)
    }

    // 4. Принимать входящие соединения
    go func() {
        for {
            conn, err := ln.Accept()
            if err != nil {
                log.Println("Ошибка соединения:", err)
                continue
            }
            mu.Lock()
            peers[conn.RemoteAddr().String()] = conn
            mu.Unlock()
            go handleConn(conn)
        }
    }()

    // 5. Чтение с клавиатуры и отправка всем peer'ам
    for {
        fmt.Print(">>> ")
        if !scanner.Scan() {
            break
        }
        text := scanner.Text()
        mu.Lock()
        for _, conn := range peers {
            fmt.Fprintf(conn, "%s\n", text)
        }
        mu.Unlock()
    }
}
