#include <stdio.h>

int main()
{
	int a;
	int b;
	int value;

	__asm
	{
		mov a, 3
		mov b, 4
		mov eax, a
		add eax, b
		mov value, eax
	}

	printf("a: %d, b: %d, value: %d\n", a, b, value);
	return 0;
}