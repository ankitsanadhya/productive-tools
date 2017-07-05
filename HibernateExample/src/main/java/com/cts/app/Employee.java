package com.cts.app;

import javax.persistence.Entity;
import javax.persistence.Id;

@Entity
public class Employee {
	
	private int employeeId;
	private String employeeName;
	
	@Id
	public int getEmployeeId() {
		return employeeId;
	}
	public void setEmployeeId(int employeeId) {
		this.employeeId = employeeId;
	}
	public String getEmployeeName() {
		return employeeName;
	}
	public void setEmployeeName(String employeeName) {
		this.employeeName = employeeName;
	}
}
