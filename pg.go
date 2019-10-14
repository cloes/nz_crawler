package main

import (
	"fmt"
	"github.com/go-pg/pg"
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
		//companyNumber, _ := strconv.Atoi(Data.CompanyNumber)

		var companyID int
		CompanyInsertSQL := "insert into company (company_number,name,nzbn,incorporation_date,company_status,entity_type,constitution_filed) values (?company_number,?company_name,?nzbn,?incorporation_date,?company_status,?entity_type,?constitution_filed) RETURNING id"
		_, err := tx.Query(pg.Scan(&companyID), CompanyInsertSQL, value)

		if err != nil {
			return err
		}

		for _, PreviousName := range Data.PreviousNames {
			PreviousNameInsertSQL := fmt.Sprintf("insert into previous_name(\"company_id\",\"name\",\"from\",\"to\") values (%d,?name,?from,?to)", companyID)
			_, err = tx.Exec(PreviousNameInsertSQL, PreviousName)

			if err != nil {
				return err
			}
		}

		for _, director := range Data.Directors {
			DirectorInsertSQL := fmt.Sprintf("insert into director (company_id,full_legal_name,residential_address,appointment_date) values (%d,?full_legal_name,?residential_address,?appointment_date)", companyID)
			_, err = tx.Exec(DirectorInsertSQL, director)

			if err != nil {
				return err
			}
		}

		var shareholdingAllocationId int
		for _, allocation := range Data.ShareholderAllocations {
			ShareholderAllocationInsertSQL := fmt.Sprintf("insert into shareholding_allocation (company_id,percentage) values (%d,?percentage) RETURNING id", companyID)
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

		AddressInsertSQL := fmt.Sprintf("insert into address (company_id,registered_office_address,address_for_service,address_for_shareregister) values (%d,?registered_office_address,?addressfor_service,?addressfor_share_register)", companyID)
		_, err = tx.Exec(AddressInsertSQL, Data)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		panic(err)
	}
}
