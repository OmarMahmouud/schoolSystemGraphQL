package main

import (
	"github.com/graphql-go/graphql"
)

func GetCoursesResolver(params graphql.ResolveParams) (interface{}, error) {
	limit, offset, err := GetRequest(params)
	if err != nil {
		return nil, err
	}
	return GetCourses(limit, offset), nil
}

func GetStudentsResolver(params graphql.ResolveParams) (interface{}, error) {
	limit, offset, err := GetRequest(params)
	if err != nil {
		return nil, err
	}
	return GetStudents(limit, offset), nil
}

func GetTeachersResolver(params graphql.ResolveParams) (interface{}, error) {
	limit, offset, err := GetRequest(params)
	if err != nil {
		return nil, err
	}
	return GetTeachers(limit, offset), nil
}

func GetStudentCourseResolver(params graphql.ResolveParams) (interface{}, error) {
	limit, offset, err := GetRequest(params)
	if err != nil {
		return nil, err
	}
	return GetStudentCourses(limit, offset), nil
}
func GetCategoriesResolver(params graphql.ResolveParams) (interface{}, error) {
	limit, offset, err := GetRequest(params)
	if err != nil {
		return nil, err
	}
	return GetCategories(limit, offset), nil
}

func GetCourseResolver(params graphql.ResolveParams) (interface{}, error) {
	id, err := getIntFromInput(params, "id", 0)
	if err != nil {
		return nil, err
	}
	return GetCourse(id), nil
}

func createStudentResolver(params graphql.ResolveParams) (interface{}, error) {
	student, err := createStudentRequest(params)
	if err != nil {
		return nil, err
	}
	return createStudent(student), nil
}

func createTeacherResolver(params graphql.ResolveParams) (interface{}, error) {
	teacher, err := createTeacherRequest(params)
	if err != nil {
		return nil, err
	}
	return createTeacher(teacher), nil
}

func createCourseResolver(params graphql.ResolveParams) (interface{}, error) {
	course, err := createCourseRequest(params)
	if err != nil {
		return nil, err
	}
	return createCourse(course), nil
}

func CreateStudentCourseResolver(params graphql.ResolveParams) (interface{}, error) {
	studentCourse, err := CreateStudentCourseRequest(params)
	if err != nil {
		return nil, err
	}
	studentCourse = GetStudentCourse(studentCourse)
	if studentCourse.ID != 0 {
		return nil, nil
	}
	return CreateStudentCourse(studentCourse), nil
}

func CreateOnLineMaterialResolver(params graphql.ResolveParams) (interface{}, error) {
	onLineMaterial, err := createOnLineMaterialRequest(params)
	if err != nil {
		return nil, err
	}

	return CreateOnLineMaterial(onLineMaterial), nil
}

func CreateOffLineMaterialResolver(params graphql.ResolveParams) (interface{}, error) {
	offLineMaterial, err := createOffLineMaterialRequest(params)
	if err != nil {
		return nil, err
	}
	return CreateOffLineMaterial(offLineMaterial), nil
}

func createCategoryResolver(params graphql.ResolveParams) (interface{}, error) {
	category, err := createCategoryRequest(params)
	if err != nil {
		return nil, err
	}
	return createCategory(category), nil
}

func updateStudentCourseResolver(params graphql.ResolveParams) (interface{}, error) {
	studentCourse, err := updateStudentCourseRequest(params)
	if err != nil {
		return nil, err
	}
	return updateStudentCourse(studentCourse), nil
}

func updateOnLineMaterialResolver(params graphql.ResolveParams) (interface{}, error) {
	onLineMaterial, err := updateOnLineMaterialRequest(params)
	if err != nil {
		return nil, err
	}
	return updateOnLineMaterial(onLineMaterial), nil
}

func updateOffLineMaterialResolver(params graphql.ResolveParams) (interface{}, error) {
	offLineMaterial, err := updateOffLineMaterialRequest(params)
	if err != nil {
		return nil, err
	}
	return updateOffLineMaterial(offLineMaterial), nil
}

func updateCourseResolver(params graphql.ResolveParams) (interface{}, error) {
	course, err := updateCourseRequest(params)
	if err != nil {
		return nil, err
	}
	return updateCourse(course), nil
}

func updateCategoryResolver(params graphql.ResolveParams) (interface{}, error) {
	category, err := updateCategoryRequest(params)
	if err != nil {
		return nil, err
	}
	return updateCategory(category), nil
}

func updateStudentResolver(params graphql.ResolveParams) (interface{}, error) {
	student, err := updateStudentRequest(params)
	if err != nil {
		return nil, err
	}
	return updateStudent(student), nil
}

func updateTeacherResolver(params graphql.ResolveParams) (interface{}, error) {
	student, err := updateTeacherRequest(params)
	if err != nil {
		return nil, err
	}
	return updateTeacher(student), nil
}

func deleteStudentResolver(params graphql.ResolveParams) (interface{}, error) {
	return deleteStudentRequest(params)
}

func deleteTeacherResolver(params graphql.ResolveParams) (interface{}, error) {
	return deleteTeacherRequest(params)
}

func deleteCourseResolver(params graphql.ResolveParams) (interface{}, error) {
	return deleteCourseRequest(params)
}

func deleteStudentCourseResolver(params graphql.ResolveParams) (interface{}, error) {
	return deleteStudentCourseRequest(params)
}

func deleteOnLineMaterialResolver(params graphql.ResolveParams) (interface{}, error) {
	return deleteOnLineMaterialRequest(params)
}

func deleteOffLineMaterialResolver(params graphql.ResolveParams) (interface{}, error) {
	return deleteOffLineMaterialRequest(params)
}

func deleteCategoryResolver(params graphql.ResolveParams) (interface{}, error) {
	return deleteCategoryRequest(params)
}
