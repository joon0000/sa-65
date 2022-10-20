import { ProvincesInterface } from "./IProvince";
import { RolesInterface } from "./IRole";
import { MemberClassesInterface } from "./IMemberClass";
import { EmployeesInterface } from "./IEmployee";
export interface UsersInterface {
    ID?:        number,
    PIN:       string,
	FirstName: string,
	LastName:  string,
	CIV:       string,
	PHONE:     string,
	Email:     string,
	Password:  string, 
	ADDRESS:   string,
	ProvinceID?:    number,
	Province?:    ProvincesInterface,  
	RoleID?:        number,    
	Role?:        RolesInterface,       
	MemberClassID?: number, 
	MemberClass?: MemberClassesInterface, 
	EmployeeID?:    number,    
	Employee?:    EmployeesInterface,    
}