package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

var schema graphql.Schema

// GraphQL endpoint
var graphqlHandler = handler.New(&handler.Config{
	Schema:   &schema,
	Pretty:   true,
	GraphiQL: true,
})

func createSchema() (graphql.Schema, error) {
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{
		Query:    graphql.NewObject(rootQuery),
		Mutation: graphql.NewObject(mutationType),
	}
	return graphql.NewSchema(schemaConfig)
}

func run() {
	r := gin.Default()
	Routing(r)
	r.Run(":8080")
}

func GetRequest(params graphql.ResolveParams) (int, int, error) {
	limit, err := getIntFromInput(params, "limit", 10)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid input type")
	}
	offset, err := getIntFromInput(params, "offset", 0)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid input type")
	}
	return limit, offset, nil
}

func getIntFromInput(params graphql.ResolveParams, key string, defaultValue int) (int, error) {
	input, ok := params.Args["input"].(map[string]interface{})
	if !ok {
		return defaultValue, fmt.Errorf("invalid input type")
	}
	value, exists := input[key]
	if !exists {
		return defaultValue, nil
	}
	return value.(int), nil
}

func createStudentRequest(params graphql.ResolveParams) (interface{}, error) {
	input, ok := params.Args["input"].(map[string]interface{})
	if !ok {
		return Student{}, fmt.Errorf("invalid input type")
	}
	student := Student{
		Name:  input["name"].(string),
		Age:   input["age"].(int),
		Email: input["email"].(string),
	}
	return student, nil
}

func createTeacherRequest(params graphql.ResolveParams) (interface{}, error) {
	input, ok := params.Args["input"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid input type")
	}
	teacher := Teacher{
		Name:     input["name"].(string),
		Age:      input["age"].(int),
		Email:    input["email"].(string),
		CourseID: uint(input["course_id"].(int)),
	}
	if input["course_id"] != nil {
		teacher.CourseID = uint(input["course_id"].(int))
	}
	return teacher, nil
}

func createCourseRequest(params graphql.ResolveParams) (interface{}, error) {
	input, ok := params.Args["input"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid input type")
	}
	course := Course{
		Name:        input["name"].(string),
		Description: input["description"].(string),
	}
	if input["category_id"] != nil {
		course.CategoryID = uint(input["category_id"].(int))
	}
	return course, nil
}

func createOnLineMaterialRequest(params graphql.ResolveParams) (interface{}, error) {
	input, ok := params.Args["input"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid input type")
	}
	onLineMaterial := OnLineMaterial{
		CourseID: uint(input["courseID"].(int)),
		Website:  input["website"].(string),
		BookName: input["bookName"].(string),
	}
	if input["price"] != nil {
		onLineMaterial.Price = uint(input["price"].(int))
	}
	return onLineMaterial, nil
}

func createOffLineMaterialRequest(params graphql.ResolveParams) (interface{}, error) {
	input, ok := params.Args["input"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid input type")
	}
	offLineMaterial := OffLineMaterial{
		CourseID: uint(input["courseID"].(int)),
		Library:  input["library"].(string),
		BookName: input["bookName"].(string),
	}
	if input["price"] != nil {
		offLineMaterial.Price = uint(input["price"].(int))
	}
	return offLineMaterial, nil
}

func updateOnLineMaterialRequest(params graphql.ResolveParams) (OnLineMaterial, error) {
	input, ok := params.Args["input"].(map[string]interface{})
	if !ok {
		return OnLineMaterial{}, fmt.Errorf("invalid input type")
	}
	onLineMaterial := OnLineMaterial{
		CourseID: uint(input["courseID"].(int)),
		Website:  input["website"].(string),
		BookName: input["bookName"].(string),
	}
	if input["price"] != nil {
		onLineMaterial.Price = uint(input["price"].(int))
	}
	return onLineMaterial, nil
}

func updateOffLineMaterialRequest(params graphql.ResolveParams) (OffLineMaterial, error) {
	input, ok := params.Args["input"].(map[string]interface{})
	if !ok {
		return OffLineMaterial{}, fmt.Errorf("invalid input type")
	}
	offLineMaterial := OffLineMaterial{
		CourseID: uint(input["courseID"].(int)),
		Library:  input["library"].(string),
		BookName: input["bookName"].(string),
	}
	if input["price"] != nil {
		offLineMaterial.Price = uint(input["price"].(int))
	}
	return offLineMaterial, nil
}

func updateCourseRequest(params graphql.ResolveParams) (Course, error) {
	input, ok := params.Args["input"].(map[string]interface{})
	if !ok {
		return Course{}, fmt.Errorf("invalid input type")
	}
	var course Course
	course.ID = uint(input["id"].(int))
	if input["name"] != nil {
		course.Name = input["name"].(string)
	}
	if input["description"] != nil {
		course.Description = input["description"].(string)
	}
	if input["category_id"] != nil {
		fmt.Println("error")
		course.CategoryID = uint(input["category_id"].(int))
	}
	return course, nil
}

func createCategoryRequest(params graphql.ResolveParams) (interface{}, error) {
	input, ok := params.Args["input"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid input type")
	}
	category := Category{
		Name:        input["name"].(string),
		Description: input["description"].(string),
	}
	return category, nil
}

func deleteStudentRequest(params graphql.ResolveParams) (interface{}, error) {
	input, ok := params.Args["input"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid input type")
	}
	deleteStudent(input["id"].(int))
	isDeleted := isDeleted{
		deleted: "yes",
	}
	return isDeleted, nil
}

func deleteTeacherRequest(params graphql.ResolveParams) (interface{}, error) {
	input, ok := params.Args["input"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid input type")
	}
	deleteTeacher(input["id"].(int))
	isDeleted := isDeleted{
		deleted: "yes",
	}
	return isDeleted, nil
}

func deleteCourseRequest(params graphql.ResolveParams) (interface{}, error) {
	input, ok := params.Args["input"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid input type")
	}
	deleteCourse(input["id"].(int))
	isDeleted := isDeleted{
		deleted: "yes",
	}
	return isDeleted, nil
}

func deleteStudentCourseRequest(params graphql.ResolveParams) (interface{}, error) {
	input, ok := params.Args["input"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid input type")
	}
	deleteStudentCourse(input["id"].(int))
	isDeleted := isDeleted{
		deleted: "yes",
	}
	return isDeleted, nil
}

func deleteOnLineMaterialRequest(params graphql.ResolveParams) (interface{}, error) {
	input, ok := params.Args["input"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid input type")
	}
	deleteOnLineMaterial(input["id"].(int))
	isDeleted := isDeleted{
		deleted: "yes",
	}
	return isDeleted, nil
}

func deleteOffLineMaterialRequest(params graphql.ResolveParams) (interface{}, error) {
	input, ok := params.Args["input"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid input type")
	}
	deleteOffLineMaterial(input["id"].(int))
	isDeleted := isDeleted{
		deleted: "yes",
	}
	return isDeleted, nil
}

func deleteCategoryRequest(params graphql.ResolveParams) (interface{}, error) {
	input, ok := params.Args["input"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid input type")
	}
	deleteCategory(input["id"].(int))
	isDeleted := isDeleted{
		deleted: "yes",
	}
	return isDeleted, nil
}

func CreateStudentCourseRequest(params graphql.ResolveParams) (StudentCourse, error) {
	input, ok := params.Args["input"].(map[string]interface{})
	if !ok {
		return StudentCourse{}, fmt.Errorf("invalid input type")
	}
	studentCourse := StudentCourse{
		CourseID:  uint(input["courseID"].(int)),
		StudentID: uint(input["studentID"].(int)),
	}
	return studentCourse, nil
}

func updateStudentCourseRequest(params graphql.ResolveParams) (interface{}, error) {
	input, ok := params.Args["input"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid input type")
	}
	id, _ := getIntFromInput(params, "id", 0)
	if id == 0 {
		return nil, fmt.Errorf("invalid input type")
	}
	studentCourse := GetStudentCourseByID(uint(id))
	if studentCourse.ID == 0 {
		return nil, fmt.Errorf("invalid input type")
	}
	if input["course_id"] != nil {
		studentCourse.CourseID = uint(input["course_id"].(int))
	}
	if input["student_id"] != nil {
		studentCourse.StudentID = uint(input["student_id"].(int))
	}
	return studentCourse, nil
}

func updateCategoryRequest(params graphql.ResolveParams) (interface{}, error) {
	input, ok := params.Args["input"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid input type")
	}
	id, _ := getIntFromInput(params, "id", 0)
	if id == 0 {
		return nil, fmt.Errorf("invalid input type")
	}
	category := GetCategoryByID(uint(id))
	if category.Name == "" {
		return nil, fmt.Errorf("invalid input type")
	}
	if input["name"] != nil {
		category.Name = input["name"].(string)
	}
	if input["description"] != nil {
		category.Description = input["description"].(string)
	}
	return category, nil
}

func updateStudentRequest(params graphql.ResolveParams) (interface{}, error) {
	input, ok := params.Args["input"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid input type")
	}
	id, _ := getIntFromInput(params, "id", 0)
	if id == 0 {
		return nil, fmt.Errorf("invalid input type")
	}
	student := GetStudentByID(uint(id))
	if student.Name == "" {
		return nil, fmt.Errorf("invalid input type")
	}
	if input["name"] != nil {
		student.Name = input["name"].(string)
	}
	if input["age"] != nil {
		student.Age = input["age"].(int)
	}
	if input["email"] != nil {
		student.Email = input["email"].(string)
	}
	return student, nil
}

func updateTeacherRequest(params graphql.ResolveParams) (interface{}, error) {
	input, ok := params.Args["input"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid input type")
	}
	id, _ := getIntFromInput(params, "id", 0)
	if id == 0 {
		return nil, fmt.Errorf("invalid input type")
	}
	teacher := GetTeacherByID(uint(id))
	if teacher.Name == "" {
		return nil, fmt.Errorf("invalid input type")
	}
	if input["name"] != nil {
		teacher.Name = input["name"].(string)
	}
	if input["age"] != nil {
		teacher.Age = input["age"].(int)
	}
	if input["email"] != nil {
		teacher.Email = input["email"].(string)
	}
	if input["course_id"] != nil {
		teacher.CourseID = input["course_id"].(uint)
	}
	return teacher, nil
}
