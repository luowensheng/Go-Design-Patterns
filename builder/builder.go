package builder

import "fmt"

type NotificationBuilder struct {
	title string
	subtitle string
	message string
	image string
	icon string
	priority int
	notType string
}

func NewNotificationBuilder() *NotificationBuilder {
	return &NotificationBuilder{}
}

func (builder *NotificationBuilder) SetTitle( title string) *NotificationBuilder{
     builder.title = title
	return builder
}
func (builder *NotificationBuilder) SetSubTitle(subtitle string) *NotificationBuilder{
	 builder.subtitle = subtitle
	return builder
}

func (builder *NotificationBuilder) SetMessage(message string) *NotificationBuilder{
	builder.message = message
	return builder
}


func (builder *NotificationBuilder) SetImage(image string) *NotificationBuilder{
	builder.image = image
	return builder
}


func (builder *NotificationBuilder) SetIcon(icon string) *NotificationBuilder{
	builder.icon = icon
	return builder
}

func (builder *NotificationBuilder) SetPriority(priority int) *NotificationBuilder{
	builder.priority = priority
	return builder
}


func (builder *NotificationBuilder) SetNotType(notType string) *NotificationBuilder{
	builder.notType = notType
	return builder
}

func (builder *NotificationBuilder) Build() (*Notification, error) {

	if builder.icon != "" && builder.subtitle == "" {
		return nil, fmt.Errorf("You need to specify a subtitle when using an icon")
	}

	if builder.priority > 5 {
		return nil, fmt.Errorf("Priority must be 0 to 5")
	}

    notification := &Notification{
		title: builder.title,
		message: builder.message,
		subtitle: builder.subtitle,
		image: builder.image,
		icon: builder.icon,
		priority: builder.priority,
		notType: builder.notType,
	} 

	return notification, nil
}

