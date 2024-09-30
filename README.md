# gorm_foreignKey

## [1] foreignKey:Device_PID:
> - Specifies that Device_PID is a foreign key that refers to the primary key (ID by default) in Devices.
> - Simple, without any special constraints.

## [2] foreignKey:Device_PID;references:ID:
> - Same as foreignKey:Device_PID, but explicitly states that the foreign key refers to the ID column. No additional constraints like cascade or set null.

## [3] foreignKey:Device_PID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL:
> - Specifies the foreign key that explicitly refers to ID, with additional constraints:
#### OnUpdate
> - : Updates the foreign key automatically when the primary key changes.
#### OnDelete
> - NULL: Sets the foreign key value to NULL if the referenced data is deleted.

## conclusion
> - case 1 & 2 : it is not possible to delete rows in the master table (Devices) if there are still relationships to other tables
> - case 3 : allows for the deletion of rows in the master table (Devices), and when the deletion occurs, the foreign key value in the referencing table (such as Foreign3) will be set to NULL.

## ref
> - [gorm](https://gorm.io/docs/belongs_to.html)

