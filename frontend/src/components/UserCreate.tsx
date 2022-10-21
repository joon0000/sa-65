import React, { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import Button from "@mui/material/Button";
import FormControl from "@mui/material/FormControl";
import Container from "@mui/material/Container";
import Paper from "@mui/material/Paper";
import Grid from "@mui/material/Grid";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import Divider from "@mui/material/Divider";
import Snackbar from "@mui/material/Snackbar";
import Select, { SelectChangeEvent } from "@mui/material/Select";
import MuiAlert, { AlertProps } from "@mui/material/Alert";
import TextField from "@mui/material/TextField";
import { AdapterDateFns } from "@mui/x-date-pickers/AdapterDateFns";
import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";
import { DatePicker } from "@mui/x-date-pickers/DatePicker";

import { ProvincesInterface } from "../interfaces/IProvince"; //Playlist --> Province
import { RolesInterface } from "../interfaces/IRole"; //Reso --> Role
import { EmployeesInterface } from "../interfaces/IEmployee"; //Video --> Employee
import { MemberClassesInterface } from "../interfaces/IMemberClass"; //Watch --> MemberClass
import { UserInterface } from "../interfaces/IUser";

import {
  GetProvinces,
  GetRoles,
  GetEmployees,
  GetMemberClasses,
  Users,
} from "../services/HttpClientService";

const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
  props,
  ref
) {
  return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});

function UserCreate() {
  const [provinces, setProvinces] = React.useState<ProvincesInterface[]>([]);
  const [roles, setRoles] = useState<RolesInterface[]>([]);
  const [memberclasses, setMemberClasses] = useState<MemberClassesInterface[]>([]);
  const [employees, setEmployees] = React.useState<EmployeesInterface[]>([]);
  const [user, setUser] = useState<Partial<UserInterface>>();

  const [success, setSuccess] = useState(false);
  const [error, setError] = useState(false);

   //เปิดปิดตัว alert
  const handleClose = (
    event?: React.SyntheticEvent | Event,
    reason?: string
  ) => {
    if (reason === "clickaway") {
      return;
    }
    setSuccess(false);
    setError(false);
  };
    //combobox
    const handleChange = (event: SelectChangeEvent) => {
    console.log(event.target.name);
    console.log(event.target.value);

    
    const name = event.target.name as keyof typeof user;
    setUser({
      ...user,
      [name]: event.target.value,
    });
  };

  //text field
  const handleInputChange = (
    event: React.ChangeEvent<{ id?: string; value: any }>
  ) => {
    const id = event.target.id as keyof typeof user;
    const { value } = event.target;
    setUser({ ...user, [id]: value });
  };

  const getProvinces = async () => {
    let res = await GetProvinces();
    if (res) {
      setProvinces(res);
    }
  };

  const getRoles = async () => {
    let res = await GetRoles();
    if (res) {
      setRoles(res);
    }
  };

  const getMemberClasses = async () => {
    let res = await GetMemberClasses();
    if (res) {
      setMemberClasses(res);
    }
  };
  
  const getEmployees = async () => {
    let res = await GetEmployees();
    if (res) {
      setEmployees(res);
    }
  };

  useEffect(() => {
    getEmployees();
    getMemberClasses();
    getProvinces();
    getRoles();
  }, []);

  const convertType = (data: string | number | undefined) => {
    let val = typeof data === "string" ? parseInt(data) : data;
    return val;
  };

  async function submit() {
    let data = {
      EmployeeID: convertType(user?.EmployeeID),
      MemberClassID: convertType(user?.MemberClassID),
      ProvinceID: convertType(user?.ProvinceID),
      RoleID: convertType(user?.RoleID),
      Pin: user?.Pin,
      FirstName: user?.FirstName,
      LastName: user?.LastName,
      Civ: user?.Civ,
      Phone: user?.Phone,
      Email: user?.Email,
      Password: user?.Password,
      Address: user?.Address,
    };

    let res = await Users(data);
    if (res) {
      setSuccess(true);
    } else {
      setError(true);
    }
  }

  return (
    <Container maxWidth="md">
      <Snackbar
        open={success}
        autoHideDuration={3000}
        onClose={handleClose}
        anchorOrigin={{ vertical: "top", horizontal: "center" }}
      >
        <Alert onClose={handleClose} severity="success">
          บันทึกข้อมูลสำเร็จ
        </Alert>
      </Snackbar>
      <Snackbar
        open={error}
        autoHideDuration={6000}
        onClose={handleClose}
        anchorOrigin={{ vertical: "top", horizontal: "center" }}
      >
        <Alert onClose={handleClose} severity="error">
          บันทึกข้อมูลไม่สำเร็จ
        </Alert>
      </Snackbar>
      <Paper>
        <Box
          display="flex"
          sx={{
            marginTop: 2,
          }}
        >
          <Box sx={{ paddingX: 2, paddingY: 1 }}>
            <Typography
              component="h2"
              variant="h6"
              color="primary"
              gutterBottom
            >
              ลงทะเบียนสมาชิก
            </Typography>
          </Box>
        </Box>
        <Divider />
        <Grid container spacing={3} sx={{ padding: 2 }}>
          
          {/*Video combobox*/}
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>พนักงานที่ลงทะเบียน</p>
              <Select
                native
                value={user?.EmployeeID + ""}
                onChange={handleChange}
                inputProps={{
                  name: "EmployeeID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกชื่อ
                </option>
                {employees.map((item: EmployeesInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Name}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>

          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>บทบาท</p>
              <Select
                native
                value={user?.RoleID + ""}
                onChange={handleChange}
                inputProps={{
                  name: "RoleID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกบทบาท
                </option>
                {roles.map((item: RolesInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Name}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>

          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>คลาส</p>
              <Select
                native
                value={user?.MemberClass + ""}
                onChange={handleChange}
                inputProps={{
                  name: "MemberClassID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกคลาส
                </option>
                {memberclasses.map((item: MemberClassesInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Name}
                </option>
                ))}
              </Select>
            </FormControl>
          </Grid>


          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>จังหวัด</p>
              <Select
                native
                value={user?.ProvinceID + ""}
                onChange={handleChange}
                inputProps={{
                  name: "ProvinceID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกจังหวัด
                </option>
                {provinces.map((item: ProvincesInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Name}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>

          <Grid item xs={12}>
            <Button
              component={RouterLink}
              to="/users"
              variant="contained"
              color="inherit"
            >
              กลับ
            </Button>
            <Button
              style={{ float: "right" }}
              onClick={submit}
              variant="contained"
              color="primary"
            >
              บันทึก
            </Button>
          </Grid>
        </Grid>
      </Paper>
    </Container>
  );
}

export default UserCreate;
