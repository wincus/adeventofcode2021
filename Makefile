ifeq ($(SESSION),)
$(error a SESSION env needs to be defined)
endif

day%:
	go run solutions/$@/main.go
