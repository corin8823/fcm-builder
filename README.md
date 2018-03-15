# fcm-builder
[![Build Status](https://travis-ci.org/corin8823/fcm-builder.svg?branch=master)](https://travis-ci.org/corin8823/fcm-builder)

topic builder of Firebase Cloud Messaging (FCM) for golang

# usage
```
// Build condition for fcm
cond, err := ToCondition(CondNot{CondTopic{"NewUserTopic"}}.And(CondTopic{"MaleUserTopic"}))
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Topic string: %d", cond)
// "Topic string: !('NewUserTopic' in topics) && 'MaleUserTopic' in topics"
```

## Acknowledgments
Inspired by [go-xorm/builder](https://github.com/go-xorm/builder)

## License
fcm-builder is available under the MIT license. See the LICENSE file for more info.
