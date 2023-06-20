package main

import (
	"gorm.io/gorm"
	"time"
)

type Course struct {
	gorm.Model
	Name            string          `json:"name" binding:"required"  gorm:"unique"`
	Description     string          `json:"description" binding:"required"`
	StartDate       time.Time       `json:"start_date"`
	EndDate         time.Time       `json:"end_date"`
	CategoryID      uint            `json:"category_id" validate:"required"`
	Category        Category        `json:"category"`
	StudentCourse   []StudentCourse `json:"student_course"`
	OnLineMaterial  OnLineMaterial
	OffLineMaterial OffLineMaterial
}

type Category struct {
	gorm.Model
	Name        string   `json:"name" validate:"required" gorm:"type:varchar(50);unique"`
	Description string   `json:"description" validate:"required"  gorm:"type:varchar(500)"`
	Courses     []Course `json:"courses"`
}

type Student struct {
	gorm.Model
	Name  string `json:"name" validate:"required"  gorm:"type:varchar(50)"`
	Age   int    `json:"age"`
	Email string `json:"email" validate:"email,required"  gorm:"unique"`
}

type StudentCourse struct {
	gorm.Model
	StudentID uint `json:"student_id"`
	CourseID  uint `json:"course_id"`
}

type isDeleted struct {
	deleted string
}

type Teacher struct {
	gorm.Model
	Name     string `json:"name" validate:"required"  gorm:"type:varchar(50)"`
	Age      int    `json:"age"`
	Email    string `json:"email" validate:"email,required"  gorm:"unique"`
	CourseID uint   `json:"course_id" gorm:"unique"`
}

type OnLineMaterial struct {
	gorm.Model
	Website  string `json:"website" validate:"required" `
	Price    uint   `json:"price"`
	BookName string `json:"book_name" validate:"required" `
	CourseID uint   `json:"course_id" validate:"required" gorm:"unique"`
}

type OffLineMaterial struct {
	gorm.Model
	Library  string `json:"library" validate:"required"`
	BookName string `json:"book_name" validate:"required" `
	Price    uint   `json:"price"`
	CourseID uint   `json:"course_id" validate:"required" gorm:"unique"`
}

func GetCourses(limit int, offset int) interface{} {
	var courses []Course
	DB.Limit(limit).Offset(offset).Preload("StudentCourse").Preload("offLineMaterialType").Preload("onLineMaterialType").Find(&courses)
	return courses
}

func GetStudents(limit int, offset int) interface{} {
	var students []Student
	DB.Limit(limit).Offset(offset).Find(&students)
	return students
}

func GetTeachers(limit int, offset int) interface{} {
	var teachers []Teacher
	DB.Limit(limit).Offset(offset).Find(&teachers)
	return teachers
}

func GetStudentCourses(limit int, offset int) interface{} {
	var studentCourse []StudentCourse
	DB.Limit(limit).Offset(offset).Find(&studentCourse)
	return studentCourse
}
func GetCategories(limit int, offset int) interface{} {
	var categories []Category
	DB.Limit(limit).Offset(offset).Preload("Courses").Find(&categories)
	return categories
}

func GetCourse(id int) interface{} {
	var course Course
	DB.Where("id = ? ", id).Find(&course)
	return course
}

func createStudent(student interface{}) interface{} {
	DB.Create(&student)
	return student
}

func createTeacher(teacher interface{}) interface{} {
	DB.Create(&teacher)
	return teacher
}

func createCourse(course interface{}) (interface{}, error) {
	DB.Create(&course)
	return course, nil
}

func CreateOnLineMaterial(onLineMaterial interface{}) (interface{}, error) {
	DB.Create(&onLineMaterial)
	return onLineMaterial, nil
}

func CreateOffLineMaterial(offLineMaterial interface{}) (interface{}, error) {
	DB.Create(&offLineMaterial)
	return offLineMaterial, nil
}

func updateOnLineMaterial(onLineMaterial OnLineMaterial) (interface{}, error) {
	DB.Where("course_id = ? ", onLineMaterial.CourseID).Updates(&onLineMaterial)
	return onLineMaterial, nil
}

func updateOffLineMaterial(offLineMaterial OffLineMaterial) (interface{}, error) {
	DB.Where("course_id = ? ", offLineMaterial.CourseID).Updates(&offLineMaterial)
	return offLineMaterial, nil
}

func updateCourse(course Course) (interface{}, error) {
	DB.Where("id = ?", course.ID).Updates(&course)
	DB.Preload("Category").Find(&course)
	return course, nil
}

func createCategory(category interface{}) (interface{}, error) {
	DB.Create(&category)
	return category, nil
}

func deleteStudent(id int) {
	DB.Where("id = ?", id).Delete(&Student{})
}

func deleteTeacher(id int) {
	DB.Where("id = ?", id).Delete(&Teacher{})
}

func deleteCourse(id int) {
	DB.Where("id = ?", id).Delete(&Course{})
}

func deleteStudentCourse(id int) {
	DB.Where("id = ?", id).Delete(&StudentCourse{})
}

func deleteOnLineMaterial(id int) {
	DB.Where("id = ?", id).Unscoped().Delete(&OnLineMaterial{})
}

func deleteOffLineMaterial(id int) {
	DB.Where("id = ?", id).Unscoped().Delete(&OffLineMaterial{})
}

func deleteCategory(id int) {
	DB.Where("id = ?", id).Delete(&Category{})
}

func GetStudentCourse(studentCourse StudentCourse) StudentCourse {
	DB.Where("student_id = ? AND course_id = ?", studentCourse.StudentID, studentCourse.CourseID).Unscoped().First(&studentCourse)
	return studentCourse
}

func CreateStudentCourse(studentCourse StudentCourse) StudentCourse {
	DB.Create(&studentCourse)
	return studentCourse
}

func GetStudentCourseByID(id uint) StudentCourse {
	var studentCourse StudentCourse
	DB.Where("id = ?", id).Find(&studentCourse)
	return studentCourse
}

func updateStudentCourse(studentCourse interface{}) interface{} {
	DB.Updates(&studentCourse)
	return studentCourse
}

func GetCategoryByID(id uint) Category {
	var category Category
	DB.Where("id = ?", id).Find(&category)
	return category
}

func updateCategory(category interface{}) interface{} {
	DB.Updates(&category)
	return category
}

func GetStudentByID(id uint) Student {
	var student Student
	DB.Where("id = ?", id).Find(&student)
	return student
}

func updateStudent(student interface{}) interface{} {
	DB.Updates(&student)
	return student
}

func GetTeacherByID(id uint) Teacher {
	var teacher Teacher
	DB.Where("id = ?", id).Find(&teacher)
	return teacher
}

func updateTeacher(teacher interface{}) interface{} {
	DB.Updates(&teacher)
	return teacher
}
