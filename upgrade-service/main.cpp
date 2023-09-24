#include <bits/stdc++.h>

using namespace std;

// Uncomment the following line if you need it
// typedef long long ll;

int main() {
	long long n, m, k,temp,sig,did=0,minsum=0;
	cin.sync_with_stdio(false);
	queue<long long> q1;
	queue<long long> q2;
	cin >> n >> m >> k;
	for (long long i = 1; i <= n; i++) {
		 long long ai;
		cin >> ai;
		q1.push(ai);    
	}
	for (long long j=1;j<=m;j++)
	{
		minsum+=q1.front();
		q2.push(q1.front());
		q1.pop();
	}
	temp=1;
	sig=1;
	while(!q1.empty())
	{
		temp++;
		did=did+q1.front()-q2.front();
		q2.push(q1.front());
		q1.pop();
		q2.pop();
		if (did<0)
		{
			sig=temp;
			minsum+=did;
			did=0;
		}
	}
	if ( minsum< m*k)
	{
		cout << sig<<' '<<sig+m-1<<endl;
	}
    else{
    	cout<<0<<' '<<0<<endl;
	}
	return 0;
}
