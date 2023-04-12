package handler

import (
	"log"
	"techTestK-Style/database"
	"techTestK-Style/model/entity"
	"techTestK-Style/model/request"

	"github.com/gofiber/fiber/v2"
)

func MemberHandlerCreate(c *fiber.Ctx) error {
	member := new(request.MemberCreateRequest)
	if err := c.BodyParser(member); err != nil {
		return err
	}

	newMember := entity.Member{
		Username: member.Username,
		Gender: member.Gender,
		SkinType: member.SkinType,
		SkinColor: member.SkinColor,
	}

	errCreateUser := database.DB.Create(&newMember).Error
	if errCreateUser != nil {
		return c.Status(500).JSON(fiber.Map{
			"message" : "failed to store data",
		})
	}
	return c.JSON(fiber.Map{
		"message" : "data successfully stored",
	})
}

func MemberHandlerUpdate(c *fiber.Ctx) error {
	memberRequest := new(request.MemberUpdateRequest)

	if err := c.BodyParser(memberRequest); err != nil {
	   return c.Status(400).JSON(fiber.Map{
		"Message" : "bad request",
	})
   }

	var member entity.Member
	
	memberId := c.Params ("id")
	
	err:= database.DB.First(&member, "ID = ?", memberId).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"Message" : "user not found",
		})
	}

	// UPDATE USER DATA
	if memberRequest.Username != "" {
		member.Username = memberRequest.Username
	}
	member.Gender = memberRequest.Gender
	member.SkinType = memberRequest.SkinType
	member.SkinColor = memberRequest.SkinColor
	

	errUpdate := database.DB.Save(&member).Error
	if errUpdate != nil {
		return c.Status(500).JSON(fiber.Map{
			"Message" : "Internal Server Error",
			"data" : member,
		})
		
	}
	return c.JSON(fiber.Map{
		"Message" : "Updated",
		"data" : member,
	})
}

func MemberHandlerDelete(c *fiber.Ctx) error {
	memberId := c.Params("id")
	var member entity.Member

	err := database.DB.Debug().First(&member, "ID = ?", memberId).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"Message": "Member not found",
		})
	}

	errDelete := database.DB.Debug().Delete(&member).Error
	if errDelete != nil {
		return c.Status(500).JSON(fiber.Map{
			"Message": "Internal Server Error",
		})
	}

	response := fiber.Map{
		"Message": "Member was deleted",
	}

	return c.JSON(response)
}

func MemberHandlerGetAll(c *fiber.Ctx) error {
	userInfo := c.Locals("userInfo")
	log.Println("user info data :: ", userInfo)

	var member[]entity.Member
	result:= database.DB.Find(&member)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return c.JSON(member)
}