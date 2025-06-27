package db

// func BeforeCreate(tx *gorm.DB) {

// 	if tx.Statement.Schema == nil {
// 		return
// 	}

// 	resource := tx.Statement.Table
// 	resourceID := tx.Statement.ReflectValue.FieldByName("ID").String()
// 	action := "CREATE"

// 	auditLog := audit.AuditLog{
// 		ResourceID:   resourceID,
// 		ResourceName: resource,
// 		Action:       action,
// 		CreatedAt:    time.Now(),
// 	}

// 	userID, ok := tx.Get("user_id")
// 	if !ok {
// 		auditLog.UserID = userID.(string)
// 	}

// 	orgID, ok := tx.Get("org_id")
// 	if !ok {
// 		auditLog.OrganizationID = orgID.(string)
// 	}

// 	if err := tx.Create(&auditLog).Error; err != nil {
// 		zap.L().Error("failed to create audit field value", zap.Error(err))
// 		return
// 	}

// 	for _, column := range tx.Statement.Schema.Fields {
// 		if column.Name != "id" {
// 			newValue := tx.Statement.ReflectValue.FieldByName(column.Name).String()

// 			auditField := audit.AuditLogFieldValue{
// 				AuditLogID:    auditLog.ID,
// 				Field:         column.Name,
// 				PreviousValue: "",
// 				NewValue:      newValue,
// 			}

// 			if err := tx.Create(&auditField).Error; err != nil {
// 				zap.L().Error("failed to create audit field value", zap.Error(err))
// 				return
// 			}
// 		}
// 	}
// }

// func BeforeSave(tx *gorm.DB) {

// 	if tx.Statement.Schema == nil {
// 		return
// 	}

// 	resource := tx.Statement.Table
// 	resourceID := tx.Statement.ReflectValue.FieldByName("ID").String()
// 	action := "UPDATE"

// 	auditLog := audit.AuditLog{
// 		ResourceID:   resourceID,
// 		ResourceName: resource,
// 		Action:       action,
// 		CreatedAt:    time.Now(),
// 	}

// 	userID, ok := tx.Get("user_id")
// 	if !ok {
// 		auditLog.UserID = userID.(string)
// 	}

// 	orgID, ok := tx.Get("org_id")
// 	if !ok {
// 		auditLog.OrganizationID = orgID.(string)
// 	}

// 	if err := tx.Create(&auditLog).Error; err != nil {
// 		zap.L().Error("failed to create audit field value", zap.Error(err))
// 		return
// 	}

// 	for _, column := range tx.Statement.Schema.Fields {
// 		if column.Name != "id" {
// 			previousValue := tx.Statement.ReflectValue.FieldByName(column.Name).String()
// 			newValue := tx.Statement.ReflectValue.FieldByName(column.Name).String()

// 			auditField := audit.AuditLogFieldValue{
// 				AuditLogID:    auditLog.ID,
// 				Field:         column.Name,
// 				PreviousValue: previousValue,
// 				NewValue:      newValue,
// 			}

// 			if err := tx.Create(&auditField).Error; err != nil {
// 				zap.L().Error("failed to create audit field value", zap.Error(err))
// 				return
// 			}
// 		}
// 	}
// }
