# Remaining

Customize the behaviour of deducting remaining uses. When some of your endpoints are more expensive than others, you can set a custom `cost` for each.


## Fields

| Field                                                                                                | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `Cost`                                                                                               | **int64*                                                                                             | :heavy_minus_sign:                                                                                   | How many tokens should be deducted from the current `remaining` value. Set it to 0, to make it free. |