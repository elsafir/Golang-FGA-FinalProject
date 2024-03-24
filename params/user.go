package params

type RegisterUser struct {
	Username string `json:"username" valid:"required~Field Username is required,maxstringlength(255)~Field Username maximum length is 255 characters,type(string)~Field Username must be string"`
	Email    string `json:"email" valid:"required~Field Email is required,maxstringlength(255)~Field Email maximum length is 255 characters,type(string)~Field Email must be string,email~Invalid Email Format"`
	Password string `json:"password" valid:"required~Field Password is required, minstringlength(6)~Password minimum length is 6 characters, maxstringlength(255)~Field Password maximum length is 255 characters,type(string)~Field Password must be string"`
	Age      int    `json:"age" valid:"required~Field Age is required,type(int)~Field Age must be number,range(8|100)~Field Age between 8 until 100"`
}

type LoginUser struct {
	Email    string `json:"email" valid:"required~Field Email is required, maxstringlength(255)~Field Email maximum length is 255 characters,type(string)~Field Email must be string,email~Invalid Email Format"`
	Password string `json:"password" valid:"required~Field Password is required, minstringlength(6)~Password minimum length is 6 characters, maxstringlength(255)~Field Password maximum length is 255 characters,type(string)~Field Password must be string"`
}

type UpdateUser struct {
	ID       uint   `json:"id" valid:"required"`
	Username string `json:"username" valid:"required~Field Username is required,maxstringlength(255)~Field Username maximum length is 255 characters,type(string)~Field Username must be string"`
	Email    string `json:"email" valid:"required~Field Email is required, maxstringlength(255)~Field Email maximum length is 255 characters,type(string)~Field Email must be string,email~Invalid Email Format"`
}
