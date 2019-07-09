package main

import (
	"fmt"
	"github.com/go-pg/pg"
	"strconv"
)

func insert(value *PageData) {
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "123456",
		Database: "nz_company",
	})
	defer db.Close()

	err := db.RunInTransaction(func(tx *pg.Tx) error {
		fmt.Println("================")
		companyNumber, _ := strconv.Atoi(Data.CompanyNumber)
		_, err := tx.Exec("insert into company (company_number,name,nzbn,incorporation_date,company_status,entity_type,constitution_filed) values (?company_number,?company_name,?nzbn,?incorporation_date,?company_status,?entity_type,?constitution_filed) RETURNING id", value)

		if err != nil {
			return err
		}

		for _, director := range Data.Directors {
			DirectorInsertSQL := fmt.Sprintf("insert into director (company_number,full_legal_name,residential_address,appointment_date) values (%d,?full_legal_name,?residential_address,?appointment_date)", companyNumber)
			_, err = tx.Exec(DirectorInsertSQL, director)

			if err != nil {
				return err
			}
		}

		var shareholdingAllocationId int
		for _, allocation := range Data.ShareholderAllocations {
			ShareholderAllocationInsertSQL := fmt.Sprintf("insert into shareholding_allocation (company_number,percentage) values (%d,?percentage) RETURNING id", companyNumber)
			_, err = tx.Query(pg.Scan(&shareholdingAllocationId), ShareholderAllocationInsertSQL, allocation)

			if err != nil {
				return err
			}

			for _, shareholder := range allocation.Shareholders {
				ShareholderInsertSQL := fmt.Sprintf("insert into shareholder (shareholding_allocation_id,name,address) values (%d,?name,?address)", shareholdingAllocationId)
				_, err := tx.Exec(ShareholderInsertSQL, shareholder)

				if err != nil {
					return err
				}
			}
		}
		return nil
	})

	if err != nil {
		panic(err)
	}
}
