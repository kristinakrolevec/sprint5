package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"
)

// создайте структуру Training
type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// создайте метод Parse()
func (t *Training) Parse(datastring string) (err error) {
	slice := strings.Split(datastring, ",")
	if len(slice) != 3 {
		err = errors.New("ошибка, данные указаны неверно")
		return err
	}
	t.Steps, err = strconv.Atoi(slice[0])

	if slice[1] != "Бег" && slice[1] != "Ходьба" {
		err = errors.New("ошибка в указании вида тренировки")
		return err
	}
	t.TrainingType = slice[1]
	t.Duration, err = time.ParseDuration(slice[2])

	return err
}

// создайте метод ActionInfo()
func (t Training) ActionInfo() string {
	resStr := ""
	switch t.TrainingType {
	case "Бег":
		run := spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Duration)
		resStr = fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), spentenergy.Distance(t.Steps), spentenergy.MeanSpeed(t.Steps, t.Duration), run)

	case "Ходьба":
		walk := spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		resStr = fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), spentenergy.Distance(t.Steps), spentenergy.MeanSpeed(t.Steps, t.Duration), walk)

	default:
		resStr = "неизвестный тип тренировки"
	}
	return resStr
}
