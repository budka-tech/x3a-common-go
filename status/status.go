package status

type Status = uint32

const (
  OK Status = iota // Успешно
  Any  // Любая ошибка
  NotFound  // Не найдено
  AlreadyExists  // Уже имеется
  PermissionDenied  // Недостаточно прав
  TooFrequentRequests  // Слишком частые запросы
  InternalError  // Внутренняя ошибка сервера
  ManyConnections  // Много подключений
  NotEnoughArguments  // Недостаточно аргументов
  NotAuthorized  // Не авторизован
  Authorized  // Уже авторизован
  IncorrectValue  // Неверное значение
  Inactivity  // Бездействие
  Timeout  // Время ожидания истекло
  ResourceUnavailable  // Ресурс недоступен
  OperationFailed  // Операция не удалась
  NotImplemented  // Не реализован
)
var m = map[Status]string{
  OK:"Успешно",
  Any:"Любая ошибка",
  NotFound:"Не найдено",
  AlreadyExists:"Уже имеется",
  PermissionDenied:"Недостаточно прав",
  TooFrequentRequests:"Слишком частые запросы",
  InternalError:"Внутренняя ошибка сервера",
  ManyConnections:"Много подключений",
  NotEnoughArguments:"Недостаточно аргументов",
  NotAuthorized:"Не авторизован",
  Authorized:"Уже авторизован",
  IncorrectValue:"Неверное значение",
  Inactivity:"Бездействие",
  Timeout:"Время ожидания истекло",
  ResourceUnavailable:"Ресурс недоступен",
  OperationFailed:"Операция не удалась",
  NotImplemented:"Не реализован",
}

func Readable(s Status) string {
  return m[s]
}