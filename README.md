# WB Задания уровня L1
### Вопросы
1. #### Какой самый эффективный способ конкатенации строк?

    Строки являются неизменяемым слайсом байт. При конкатенации двух строк происходит выделение новой памяти.
    Это может негативно сказать на производительности программы и на потреблении памяти, если необходимо соединять большое количество строк.
    Поэтому стоит использовать `strings.Builder` 

    ```go
    var b strings.Builder
    for i := 0; i < 5; i++ {
        b.WriteString("some-text")
   }
    fmt.Println(b.String())
    ```

2. #### Что такое интерфейсы, как они применяются в Go?
    `Интерфейсы` — это инструменты для определения наборов действий и поведения.
    Они позволяют объектам опираться на абстракции, а не фактические реализации других объектов. 
    
    При этом для компоновки различных поведений можно группировать несколько интерфейсов. 
    С помощью `интерфейсов` можно организовывать разные группы методов, применяемых к разным объектам. 
    
    Таким образом, программа вместо фактических реализаций сможет опираться на более высокие абстракции (интерфейсы), позволяя методам работать с различными объектами, реализующими один и тот же интерфейс. 
    Этот принцип называется инверсией зависимостей.

    В Go можно автоматически сделать вывод, что `структура (объект)` реализует `интерфейс`, когда она реализуется все его методы.
    
   ```go
    type Printer interface {
        Print()  
    }
    ```
   
   Это очень простой `интерфейс`, который определяет метод `Print()`. Данный метод представляет действие или поведение, которые могут реализовывать другие объекты.

   `Интерфейсы` определяют только поведение, но не фактические реализации. Это уже работа объекта, реализующего данный интерфейс.

    ##### Как устроен Duck-typing в Go?

    ```
    Если это выглядит как утка, плавает как утка и крякает как утка, то это, вероятно, утка и есть.
   ```
    
    Если структура содержит в себе все методы, что объявлены в интерфейсе, и их сигнатуры совпадают — она автоматически удовлетворяет интерфейс.

    Такой подход позволяет полиморфно (полиморфизм — способность функции обрабатывать данные разных типов) работать с объектами, которые не связаны в иерархии наследования. Достаточно, чтобы все эти объекты поддерживали необходимый набор методов.

    ##### Пустой interface{}
    Ему удовлетворяет вообще любой тип. Пустой интерфейс ничего не означает, никакой абстракции. Поэтому использовать пустые интерфейсы нужно в самых крайних случаях.
    
    ##### Интерфейсный тип

    В Go интерфейсный тип выглядит вот так:

    ```go
    type iface struct {
        tab  *itab
        data unsafe.Pointer
    }
    ```
    Где `tab` — это указатель на `Interface Table` или `itable` — структуру, которая хранит некоторые метаданные о типе и список методов, используемых для удовлетворения интерфейса, а `data` указывает на реальную область памяти, в которой лежат данные изначального объекта (статическим типом).

    Компилятор генерирует метаданные для каждого статического типа, в которых, помимо прочего, хранится список методов, реализованных для данного типа. Аналогично генерируются метаданные со списком методов для каждого интерфейса. Теперь, во время исполнения программы, runtime Go может вычислить `itable` на лету (late binding) для каждой конкретной пары. Этот `itable` кешируется, поэтому просчёт происходит только один раз.

    Зная это, становится очевидно, почему Go ловит несоответствия типов на этапе компиляции, но кастинг к интерфейсу — во время исполнения.

    Что важно помнить — переменная интерфейсного типа может принимать `nil`. Но так как объект интерфейса в Go содержит два поля: `tab` и `data` — по правилам Go, интерфейс может быть равен `nil` только если оба этих поля не определены.

3. #### Чем отличаются `RWMutex` от `Mutex`?
    `Mutex` блокирует доступ к общим ресурсам методами `.Lock()` и `.Unlock()`. 
    У `RMutex` есть дополнительные методы `.RLock()` и `.RUnlock()`, они позволяют совершать параллельное чтение общего ресурса.
    Т.е. при чтении другие вызовы `.RLock()` и `.RUnlock()` не будут заблокированы, а вот вызовы `.Lock()` и `.Unlock()` будут заблокированы.
4. #### Чем отличаются буферизированные и не буферизированные каналы?
    Основное отличие заключается в емкости: у небуферизированного канала она составляет 1, а у буферизированного - N.

    Чтение или запись данных в небуферизированный канал блокирует горутину и контроль передается свободной горутине.

    Буферизированный канал создается указанием размера буфера, в этом случае горутина не блокируется до тех пор, пока буфер не будет заполнен.

    ```go
    //небуферизированный канал
    ch1 := make(chan struct{})
    ch1 <- struct{}{} //ок, ушло в канал
    ch1 <- struct{}{} //лок. ждем когда кто-нибудь прочитает из канала
    
    //буферизированный канал
    ch2 := make(chan struct{},3)
    ch2 <- struct{}{} //ок, ушло в канал
    ch2 <- struct{}{} //ушло
    ch2 <- struct{}{} //ушло
    ch2 <- struct{}{} //тут лок. ждем когда кто-то прочитает из канала
    ```
5. #### Какой размер у структуры `struct{}{}`?
    Пустые структуры занимают 0 байт

6. #### Есть ли в Go перегрузка методов или операторов?
    Нет

7. #### В какой последовательности будут выведены элементы `map[int]int`?
    Пример:
    ```go
    m[0]=1
    m[1]=124
    m[2]=281
    ```
    Последовательность вывода будет случайной, тк порядок последовательности итераций по `map` не гарантирована.

8. #### В чем разница `make` и `new`?

    Встроенная функция `new` выделяет память (создает неименованную переменную и возваращет указатель на ее значение). 

    Аргумент - тип, и возвращаемое значение - указатель на нулевое значение указанного типа.
    ```go
    p := new(int)   // p has *int type 
    fmt.Println(*p) // "0"
    ```
   Функция `make` — это специальная встроенная функция, которая используется для инициализации слайсов, мап и каналов.
   `make` можно использовать только для инициализации слайсов, мап и каналов, и что, в отличие от функции `new`, `make` не возвращает указатель.
   
    ```go
    s := make([]int, 2,3)   //слайс
    m := make(map[int]int)  //мапа
    ch := make(chan int)    //небуфер-й канал
    ch := make(chan int, 2) //буфер-й канал
    ```
    
9. #### Сколько существует способов задать переменную типа `slice` или `map`?
    `slice` - 5

    `map` - 4
    ```go
    //5 способов задать slice
    slice := make([]int, 0)	
    slice := make([]int, 0, 2)	
    var slice []int	
    slice := []int{}	
    sice := []int{1,2}
        
    //4 способа задать map	
    var m map[string]int	
    m := map[string]int{		
        "1": 1,
        "3": 20,
    }	
    m := make(map[string]int)	
    m := make(map[string]int, 10)
    ```

10. #### Что выведет данная программа и почему?

    ```go
    func update(p *int) {
        b := 2
        p = &b
    }
    
    func main() {
        var (
            a = 1
            p = &a
        )
        fmt.Println(*p)
        update(p)
        fmt.Println(*p)
    }
    ```
    Программа выведет:
    ```go
    1
    1
    ```
    Несмотря на то, что в функцию `update` передается указатель, переменная в главной функции не изменится. 
    Обе переменные `p` указывают на один адрес, но является разными => переменная `p` в главной функции не изменилась.
    
    Чтобы это исправить необходимо возвращать из функции `update` переменную `p` и в главной функции присваивать это значение переменной

11. #### Что выведет данная программа и почему?

    ```go
    func main() {
        wg := sync.WaitGroup{}
        for i := 0; i < 5; i++ {
            wg.Add(1)
            go func(wg sync.WaitGroup, i int) {
                fmt.Println(i)
                wg.Done()
            }(wg, i)
        }
        wg.Wait()
        fmt.Println("exit")
    }
    ```
    Будет deadlock. Так как в анонимную функцию мы передаем `WitGroup` по значению, то при вызове `wg.Done()` счетчик не уменьшит значение `WaitGroup` глобально. Поэтому `main()` горутина навсегда зависает на `wg.Wait()`

    Чтобы это исправить необходимо передавать в фунцию `WaitGroup` по указателю, а не значению

12. #### Что выведет данная программа и почему?

    ```go
    func main() {
        n := 0
        if true {
            n := 1
            n++
        }
        fmt.Println(n)
    }
    ```
    Программа выведет `0`. 

    Так как в конструкции условия `if true{...}` переменная `n` локально переопределяется  
    и инкрементируется также локально. Но выводится в консоли значение переменной `n`, которая находится в другой области видимости, вне конструкции условия.

13. #### Что выведет данная программа и почему?

    ```go
    func someAction(v []int8, b int8) {
        v[0] = 100
        v = append(v, b)
    }
    
    func main() {
        var a = []int8{1, 2, 3, 4, 5}
        someAction(a, 6)
        fmt.Println(a)
    }
    ```
    Программа выведет `[100 2 3 4 5]`
    
    Слайс является структурой, содержащей в себе длину, емкость и указатель на массив.
    Передавая слайс по значению, а не по указателю, то в функции он будет копироваться. Поэтому при изменении слайса в функции `someAction`  изменения никак не отразятся на слайсе `a` в функции `main`
14. #### Что выведет данная программа и почему?

    ```go
    func main() {
        slice := []string{"a", "a"}
    
        func(slice []string) {
            slice = append(slice, "a")
            slice[0] = "b"
            slice[1] = "b"
            fmt.Print(slice)
        }(slice)
        fmt.Print(slice)
    }
    ```
    Программа выведет `[b b a] [a a]`
    
    Так как слайс передается в функцию по значению, то он копируется. Изменения слайса будут только внутри функции, за ее пределами(в данном случае в функции `main()`) слайс останется прежним.
    

