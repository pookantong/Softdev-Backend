package service

import (
	"errors"
	"fmt"

	"github.com/pkstpm/Softdev-Backend/internal/restaurant/dto"
	"github.com/pkstpm/Softdev-Backend/internal/restaurant/model"
	"github.com/pkstpm/Softdev-Backend/internal/restaurant/repository"
)

type restaurantServiceImpl struct {
	restaurantRepository repository.RestaurantRepository
}

func NewRestaurantService(restaurantRepository repository.RestaurantRepository) RestaurantService {
	return &restaurantServiceImpl{restaurantRepository: restaurantRepository}
}

func (r *restaurantServiceImpl) FindRestaurantByName(name string) ([]model.Restaurant, error) {
	restaurants, err := r.restaurantRepository.FindRestaurantByName(name)
	if err != nil {
		return nil, err
	}
	return restaurants, nil
}

func (r *restaurantServiceImpl) FindRestaurantByCategory(category string) ([]model.Restaurant, error) {
	restaurants, err := r.restaurantRepository.FindRestaurantByCategory(category)
	if err != nil {
		return nil, err
	}
	return restaurants, nil
}

func (r *restaurantServiceImpl) CreateDish(userId string, dto *dto.CreateDishDTO, imgPath string) error {
	restaurant, err := r.restaurantRepository.FindRestaurantByUserID(userId)
	if err != nil {
		return err
	}

	_, err = r.restaurantRepository.FindDishByName(dto.Name, restaurant.ID.String())
	if err == nil {
		return errors.New("dish name already exists")
	}

	var dish = &model.Dish{
		RestaurantID: restaurant.ID,
		Name:         dto.Name,
		Description:  dto.Description,
		Price:        dto.Price,
		ImgPath:      imgPath,
	}

	err = r.restaurantRepository.CreateDish(dish)
	if err != nil {
		return err
	}
	return nil
}

func (r *restaurantServiceImpl) UpdateDish(userId string, dto *dto.UpdateDishDTO) error {
	restaurant, err := r.restaurantRepository.FindRestaurantByUserID(userId)
	if err != nil {
		return err
	}

	dish, err := r.restaurantRepository.FindDishById(dto.ID)
	if err != nil {
		return err
	}

	if restaurant.ID != dish.RestaurantID {
		return errors.New("dish does not belong to restaurant")
	}

	_, err = r.restaurantRepository.FindDishByName(dto.Name, restaurant.ID.String())
	if err == nil {
		return errors.New("dish name already exists")
	}

	dish.Name = dto.Name
	dish.Description = dto.Description
	dish.Price = dto.Price

	err = r.restaurantRepository.UpdateDish(dish)
	if err != nil {
		return err
	}
	return nil
}

func (r *restaurantServiceImpl) GetAllRestaurants() ([]model.Restaurant, error) {
	restaurants, err := r.restaurantRepository.GetAllRestaurants()
	if err != nil {
		return nil, err
	}
	return restaurants, nil
}

func (r *restaurantServiceImpl) CreateTimeSlot(userId string) error {

	restaurant, err := r.restaurantRepository.FindRestaurantByUserID(userId)
	if err != nil {
		return err
	}

	for i := 0; i <= 6; i++ {
		timeslot := &model.TimeSlot{
			RestaurantID: restaurant.ID,
			Weekday:      i,
			HourStart:    9,
			HourEnd:      21,
			IsClosed:     false,
		}
		err = r.restaurantRepository.CreateTimeSlot(timeslot)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *restaurantServiceImpl) GetTimeSlot(userId string) ([]model.TimeSlot, error) {
	restaurant, err := r.restaurantRepository.FindRestaurantByUserID(userId)
	if err != nil {
		return nil, err
	}
	timeSlots, err := r.restaurantRepository.GetTimeSlotsByRestaurantId(restaurant.ID.String())
	if err != nil {
		return nil, err
	}
	fmt.Println(timeSlots)
	return timeSlots, err
}

func (r *restaurantServiceImpl) GetTimeSlotByRestaurantId(restaurantId string) ([]model.TimeSlot, error) {
	timeSlots, err := r.restaurantRepository.GetTimeSlotsByRestaurantId(restaurantId)
	if err != nil {
		return nil, err
	}
	return timeSlots, nil
}

func (r *restaurantServiceImpl) UpdateTimeSlot(userId string, dto *dto.UpdateTimeDTO) error {
	restaurant, err := r.restaurantRepository.FindRestaurantByUserID(userId)
	if err != nil {
		return err
	}

	timeslots, err := r.restaurantRepository.GetTimeSlotsByRestaurantId(restaurant.ID.String())

	if err != nil {
		return err
	}

	for _, timeslot := range timeslots {
		for _, dto := range dto.TimeSlots {
			if timeslot.Weekday == dto.Weekday {
				timeslot.HourStart = dto.HourStart
				timeslot.HourEnd = dto.HourEnd
				timeslot.IsClosed = dto.IsClosed
				timeslot.Slots = dto.Slots
				err = r.restaurantRepository.UpdateTimeSlot(&timeslot)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (r *restaurantServiceImpl) GetTablesByRestaurantId(restaurantId string) ([]model.Table, error) {
	tables, err := r.restaurantRepository.GetTablesByRestaurantId(restaurantId)
	if err != nil {
		return nil, err
	}
	return tables, nil
}

func (r *restaurantServiceImpl) GetAllDishesByRestaurantId(restaurantId string) ([]model.Dish, error) {
	dishes, err := r.restaurantRepository.GetAllDishesByRestaurantId(restaurantId)
	if err != nil {
		return nil, err
	}
	return dishes, nil
}

func (r *restaurantServiceImpl) CreateTable(userId string, dto *dto.CreateTableDTO) error {
	restaurant, err := r.restaurantRepository.FindRestaurantByUserID(userId)
	if err != nil {
		return err
	}

	table := &model.Table{
		RestaurantID: restaurant.ID,
		TableNumber:  dto.TableNumber,
		Capacity:     dto.Capacity,
	}

	err = r.restaurantRepository.CreateTable(table)
	if err != nil {
		return err
	}

	return nil
}

func (r *restaurantServiceImpl) UploadRestaurantPictures(userId string, uploadedFiles []string) error {
	restaurant, err := r.restaurantRepository.FindRestaurantByUserID(userId)
	if err != nil {
		return err
	}

	for _, file := range uploadedFiles {
		err = r.restaurantRepository.CreateImages(&model.Image{
			RestaurantID: restaurant.ID,
			ImgPath:      file,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *restaurantServiceImpl) GetRestaurantByID(restaurantId string) (*model.Restaurant, error) {
	restaurant, err := r.restaurantRepository.FindRestaurantByID(restaurantId)
	if err != nil {
		return nil, err
	}
	return restaurant, nil
}

func (r *restaurantServiceImpl) DeletetRestaurantPicture(userId string, pictureId string) error {
	restaurant, err := r.restaurantRepository.FindRestaurantByUserID(userId)
	if err != nil {
		return err
	}

	err = r.restaurantRepository.DeleteImage(restaurant.ID.String())
	if err != nil {
		return err
	}

	return nil
}

func (r *restaurantServiceImpl) GetRestaurantByUserId(userId string) (*model.Restaurant, error) {
	restaurant, err := r.restaurantRepository.FindRestaurantByUserID(userId)
	if err != nil {
		return nil, err
	}
	return restaurant, nil
}
