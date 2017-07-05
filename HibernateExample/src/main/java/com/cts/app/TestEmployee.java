package com.cts.app;

import org.hibernate.Session;
import org.hibernate.SessionFactory;
import org.hibernate.cfg.Configuration;
import org.hibernate.tool.hbm2ddl.SchemaExport;

public class TestEmployee {

	public static void main(String[] args) {
		Configuration configuration = new Configuration();
		configuration.addAnnotatedClass(Employee.class);
		configuration.configure();
		
		new SchemaExport(configuration).create(true, true);
		//SessionFactory sessionFactory = configuration.configure().buildSessionFactory();
		//Session session = sessionFactory.openSession();
		
		//2 Scenario 
		//Employee employee = new Employee();
		//employee.setEmployeeId(1);
		//employee.setEmployeeName("ABC");
		//session.save(employee);
	}

}
